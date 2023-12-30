package main

import (
	"fmt"
	"path/filepath"
	"pusat-rumah-lelang-backend/db/seeder"
	"pusat-rumah-lelang-backend/internal/config"
	"pusat-rumah-lelang-backend/internal/helper"
	"pusat-rumah-lelang-backend/internal/migrations"
	"pusat-rumah-lelang-backend/internal/model"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

func getValueFromRow(row []string, index int) string {
	if len(row) > index {
		return strings.Trim(row[index], "")
	}
	return ""
}

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	minio, err := config.NewMinioStorage(viperConfig)
	if err != nil {
		panic(err.Error())
	}

	photoPaths, err := helper.RetrieveFiles(filepath.Join("", "data/photo"))
	if err != nil {
		panic(err.Error())
	}

	migrations.DropTable(db)
	migrations.RunMigrations(db)

	seeder.Seeds(db)

	filePath := "data/prl-list-rumah.xlsx"
	// Open the Excel file
	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all the rows from a specific sheet (assume it's the first sheet, index 1)
	rows, err := xlsx.GetRows(xlsx.GetSheetName(1))
	if err != nil {
		fmt.Println(err)
		return
	}

	// skip first 4 rows
	rows = rows[3:]
	// Process each row

	if err := processWithData(db, rows, photoPaths, minio); err != nil {
		fmt.Println("Error processing data with transaction:", err)
		return
	}

	fmt.Println("Data processed successfully with transaction!")
}

func upload(minio *helper.MinioStorage, key string, index int, photoPath string) (filename string, err error) {
	filename, err = minio.UploadFile(fmt.Sprintf("%s/%s/%d", "photo_house", key, index), filepath.Join("RootDir", "/data/photo", photoPath))
	return
}

func processWithData(db *gorm.DB, rows [][]string, photoPath map[string][]string, minio *helper.MinioStorage) error {
	for _, row := range rows {
		// no := getValueFromRow(row, 0)
		owner := getValueFromRow(row, 1)
		bank := getValueFromRow(row, 2)
		// photo_house := getValueFromRow(row, 3)
		title := getValueFromRow(row, 4)
		price := getValueFromRow(row, 5)
		address := getValueFromRow(row, 6)
		latitude := getValueFromRow(row, 7)
		longitude := getValueFromRow(row, 8)
		bedroom := getValueFromRow(row, 9)
		road_access := getValueFromRow(row, 10)
		water_source := getValueFromRow(row, 11)
		area := getValueFromRow(row, 12)
		var areas []string
		var landArea string
		var buildingArea string

		if area != "" {
			areas = strings.Split(area, "/")
			landArea = strings.Trim(areas[0], " ")
			buildingArea = strings.Trim(areas[1], " ")
		}
		electricity_capacity := getValueFromRow(row, 13)
		selling_status := getValueFromRow(row, 14)
		certificate := getValueFromRow(row, 15)
		description := getValueFromRow(row, 16)

		priceFloat, err := strconv.ParseFloat(helper.RemoveCommas(price), 64)
		if err != nil {
			panic(err)
		}

		var bedroomInt int
		if bedroom != "" {
			bedroomInt, err = strconv.Atoi(bedroom)
			if err != nil {
				panic(err)
			}
		}

		// Begin transaction block for database operations
		if err := db.Transaction(func(tx *gorm.DB) error {
			certificates := &model.Certificate{}
			if err := tx.FirstOrCreate(&certificates, model.Certificate{Name: certificate, Description: certificate}).Error; err != nil {
				return err
			}

			banks := &model.Bank{}
			if err := tx.FirstOrCreate(&banks, model.Bank{Name: bank, Description: bank}).Error; err != nil {
				return err
			}

			sellingStatus := &model.SellingStatus{}
			if err = tx.FirstOrCreate(&sellingStatus, model.SellingStatus{Name: selling_status}).Error; err != nil {
				return err
			}

			roadAcess := model.RoadAccess{}
			if err = tx.FirstOrCreate(&roadAcess, model.RoadAccess{Name: road_access, Description: road_access}).Error; err != nil {
				return err
			}

			var sellingStatusID *uint
			if selling_status != "" {
				sellingStatusID = &sellingStatus.ID
			}

			var certificateID *uint
			if selling_status != "" {
				certificateID = &certificates.ID
			}

			property := &model.Property{}
			if err = tx.Where(model.Property{Title: title, BankID: &banks.ID}).Assign(model.Property{
				Title:               title,
				Owner:               owner,
				Price:               priceFloat,
				Address:             address,
				Latitude:            latitude,
				Longitude:           longitude,
				Bedrooms:            bedroomInt,
				RoadAccessID:        &roadAcess.ID,
				BankID:              &banks.ID,
				WaterSource:         water_source,
				LandArea:            landArea,
				BuildingArea:        buildingArea,
				ElectricityCapacity: electricity_capacity,
				SellingStatusID:     sellingStatusID,
				Description:         description,
				CertificateID:       certificateID,
			}).FirstOrCreate(&property).Error; err != nil {
				return err
			}

			for index, photoUrl := range photoPath[title] {
				photoHouse := &model.PhotoHouse{}
				nameFile, err := upload(minio, title, index, photoUrl)
				if err != nil {
					panic(err.Error())
				}
				tx.FirstOrCreate(&photoHouse, model.PhotoHouse{
					PropertyID: property.ID,
					PhotoUrl:   nameFile,
				})
			}
			return nil
		}); err != nil {
			return err
		}
	}
	return nil
}

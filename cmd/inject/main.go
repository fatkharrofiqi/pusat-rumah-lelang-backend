package main

import (
	"fmt"
	"pusat-rumah-lelang-backend/config"
	"pusat-rumah-lelang-backend/helpers"
	"pusat-rumah-lelang-backend/models"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

func getValueFromRow(row []string, index int) string {
	if len(row) > index {
		return row[index]
	}
	return ""
}

func main() {
	config.LoadEnv()
	config.LoadConstant()
	db := config.OpenDB()

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

	if err := processWithData(db, rows); err != nil {
		fmt.Println("Error processing data with transaction:", err)
		return
	}

	fmt.Println("Data processed successfully with transaction!")
}

func processWithData(db *gorm.DB, rows [][]string) error {
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
			landArea = areas[0]
			buildingArea = areas[1]
		}
		electricity_capacity := getValueFromRow(row, 13)
		selling_status := getValueFromRow(row, 14)
		description := getValueFromRow(row, 15)

		priceFloat, err := strconv.ParseFloat(helpers.RemoveCommas(price), 64)
		if err != nil {
			panic(err)
		}

		bedroomInt, err := strconv.Atoi(bedroom)
		if err != nil {
			panic(err)
		}

		// Begin transaction block for database operations
		if err := db.Transaction(func(tx *gorm.DB) error {
			banks := &models.Bank{}
			if err := tx.FirstOrCreate(&banks, models.Bank{Name: bank, Description: bank}).Error; err != nil {
				return err
			}

			sellingStatus := &models.SellingStatus{}
			if err = tx.FirstOrCreate(&sellingStatus, models.SellingStatus{Name: selling_status}).Error; err != nil {
				return err
			}

			roadAcess := models.RoadAccess{}
			if err = tx.FirstOrCreate(&roadAcess, models.RoadAccess{Name: road_access, Description: road_access}).Error; err != nil {
				return err
			}

			if err = tx.Where(models.Property{Title: title, BankID: &banks.ID}).Assign(models.Property{
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
				SellingStatusID:     &sellingStatus.ID,
				Description:         description,
			}).FirstOrCreate(&models.Property{}).Error; err != nil {
				return err
			}

			return nil
		}); err != nil {
			return err
		}
	}
	return nil
}

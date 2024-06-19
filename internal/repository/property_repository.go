package repository

import (
	"errors"
	"pusat-rumah-lelang-backend/internal/model"
	"pusat-rumah-lelang-backend/internal/request"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IPropertyRepository interface {
	GetTotalByCategory(db *gorm.DB, category string) (result []*model.NameCount, err error)
	Create(db *gorm.DB, request *request.CreatePropertyRequest) error
	Update(db *gorm.DB, id int64, request *request.UpdatePropertyRequest) error
	Delete(db *gorm.DB, id int64) error
	GetAll(db *gorm.DB, request *request.GetAllPropertyRequest) (result []*model.Property, total int64, err error)
	GetById(db *gorm.DB, id int64) (result *model.Property, err error)
	GetBySellingStatus(db *gorm.DB, id int64, request *request.GetBySellingStatusRequest) (result []*model.Property, err error)
	GetByLocation(db *gorm.DB, request *request.GetByLocationRequest) (result []*model.Property, err error)
}

type PropertyRepository struct {
	Repository[model.Property]
	Log *logrus.Logger
}

func NewPropertyRepository(log *logrus.Logger) IPropertyRepository {
	return &PropertyRepository{Log: log}
}

func (r *PropertyRepository) GetTotalByCategory(db *gorm.DB, category string) (result []*model.NameCount, err error) {
	var dynamicField string

	switch category {
	case "selling_status":
		dynamicField = "selling_statuses"
	case "bank":
		dynamicField = "banks"
	default:
		return nil, errors.New("invalid category")
	}

	if err := db.Table(dynamicField).
		Select(dynamicField + ".name, " + dynamicField + ".id, COUNT(properties.id) as property_count").
		Joins("LEFT JOIN properties ON " + dynamicField + ".id = properties." + category + "_id").
		Group(dynamicField + ".name").
		Scan(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r *PropertyRepository) Create(db *gorm.DB, property *request.CreatePropertyRequest) error {
	return db.Create(property).Error
}

func (r *PropertyRepository) Update(db *gorm.DB, id int64, property *request.UpdatePropertyRequest) error {
	return db.Where("id = ?", id).Save(property).Error
}

func (r *PropertyRepository) Delete(db *gorm.DB, id int64) error {
	data := &model.Property{}
	return db.Where("id = ?", id).Delete(data).Error
}

func (r *PropertyRepository) GetAll(db *gorm.DB, request *request.GetAllPropertyRequest) (properties []*model.Property, total int64, err error) {
	offset := (request.Page - 1) * request.Size

	query := db.Table("properties").
		Joins("JOIN banks ON properties.bank_id = banks.id").
		Joins("JOIN selling_statuses ON properties.selling_status_id = selling_statuses.id").
		Joins("JOIN road_accesses ON properties.road_access_id = road_accesses.id").
		Joins("JOIN certificates ON properties.certificate_id = certificates.id")

	// Apply dynamic query conditions
	if request.Query != "" {
		fields := []string{
			"title", "owner", "address", "building_area", "land_area", "latitude", "longitude",
			"property_tax_photo", "electricity_capacity", "water_source", "properties.description",
			"banks.name", "selling_statuses.name", "road_accesses.name", "certificates.name",
		}

		var conditions []string
		var values []interface{}

		for _, field := range fields {
			conditions = append(conditions, field+" LIKE ?")
			values = append(values, "%"+request.Query+"%")
		}

		query = query.Where(strings.Join(conditions, " OR "), values...)
	}

	// Retrieve the total count based on the same dynamic conditions
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination to the data retrieva
	if err := query.Preload("SellingStatus").
		Preload("Bank").
		Preload("PhotoHouse").
		Preload("PhotoCertificate").
		Preload("RoadAccess").
		Preload("Certificate").
		Limit(request.Size).
		Offset(offset).
		Find(&properties).Error; err != nil {
		return nil, 0, err
	}

	return properties, total, nil
}

func (r *PropertyRepository) GetById(db *gorm.DB, id int64) (property *model.Property, err error) {
	if err := db.
		Preload("SellingStatus").
		Preload("PhotoHouse").
		Preload("PhotoCertificate").
		Preload("RoadAccess").
		Preload("Certificate").
		Preload("Bank").
		First(&property, id).Error; err != nil {
		r.Log.WithError(err).Error("failed to get property by id")
		return property, err
	}

	return property, nil
}

func (r *PropertyRepository) GetBySellingStatus(db *gorm.DB, id int64, request *request.GetBySellingStatusRequest) (property []*model.Property, err error) {
	offset := (request.Page - 1) * request.Size
	query := db
	if request.Latitude != "" && request.Longitude != "" && request.Radius != "" {
		sql := `SELECT
							*,
							(
								6371 *
								acos(
									cos(radians(?)) * cos(radians(latitude)) * cos(radians(longitude) - radians(?)) +
										sin(radians(?)) * sin(radians(latitude))
								)
							) as distance
						FROM
							properties p
						WHERE 
							p.selling_status_id = ?
						HAVING
							distance <= ?
						ORDER BY
							distance asc
						`
		query = db.Raw(sql, request.Latitude, request.Longitude, request.Latitude, id, request.Radius)
	}
	if err := query.
		Where("selling_status_id = ?", id).
		Preload("SellingStatus").
		Preload("PhotoHouse").
		Preload("PhotoCertificate").
		Preload("RoadAccess").
		Preload("Bank").
		Preload("Certificate").
		Limit(request.Size).
		Offset(offset).
		Find(&property).Error; err != nil {
		r.Log.WithError(err).Error("failed to get property by selling status")
		return property, err
	}

	return property, nil
}

func (r *PropertyRepository) GetByLocation(db *gorm.DB, request *request.GetByLocationRequest) (properties []*model.Property, err error) {
	offset := (request.Page - 1) * request.Size
	sql := `SELECT
						*,
						(
							6371 *
							acos(
								cos(radians(?)) * cos(radians(latitude)) * cos(radians(longitude) - radians(?)) +
									sin(radians(?)) * sin(radians(latitude))
							)
						) as distance
					FROM
						properties p
					HAVING
						distance <= ?
					ORDER BY
						distance asc
					`
	if err := db.Raw(sql, request.Latitude, request.Longitude, request.Latitude, request.Radius).
		Preload("Bank").
		Preload("SellingStatus").
		Preload("PhotoHouse").
		Preload("PhotoCertificate").
		Preload("RoadAccess").
		Preload("Certificate").
		Limit(request.Size).
		Offset(offset).
		Find(&properties).Error; err != nil {
		r.Log.WithError(err).Error("failed to get properties by location")
		return properties, err
	}

	return properties, nil
}

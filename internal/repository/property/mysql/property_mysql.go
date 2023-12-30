package mysql

import (
	"errors"
	"pusat-rumah-lelang-backend/internal/common/interfaces"
	"pusat-rumah-lelang-backend/internal/model"
	"pusat-rumah-lelang-backend/internal/request"
	"strings"

	"gorm.io/gorm"
)

type IPropertyMysql interface {
	interfaces.IGetByIdGeneric[model.Property]
	interfaces.ICreateGeneric[model.Property]
	interfaces.IDeleteGeneric[model.Property]
	interfaces.IUpdateGeneric[model.Property]
	GetAll(req request.PropertyPaginationRequest) ([]model.Property, error)
	GetBySellingStatus(id int64, page, pageSize int) (property []model.Property, err error)
	GetByLocation(latitude string, longitude string, radius string) ([]model.Property, error)
	GetTotalByCategory(category string) (result []model.NameCount, err error)
}

type PropertyMysql struct {
	db *gorm.DB
}

func NewPropertyMysql(db *gorm.DB) IPropertyMysql {
	return &PropertyMysql{db: db}
}

func (r *PropertyMysql) GetTotalByCategory(category string) (result []model.NameCount, err error) {
	var dynamicField string

	switch category {
	case "selling_status":
		dynamicField = "selling_statuses"
	case "bank":
		dynamicField = "banks"
	default:
		return nil, errors.New("invalid category")
	}

	if err := r.db.Table(dynamicField).
		Select(dynamicField + ".name, " + dynamicField + ".id, COUNT(properties.id) as property_count").
		Joins("LEFT JOIN properties ON " + dynamicField + ".id = properties." + category + "_id").
		Group(dynamicField + ".name").
		Scan(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *PropertyMysql) Create(property *model.Property) error {
	return r.db.Create(property).Error
}

func (r *PropertyMysql) Update(id int64, property *model.Property) error {
	return r.db.Where("id = ?", id).Save(property).Error
}

func (r *PropertyMysql) Delete(id int64) error {
	data := &model.Property{}
	return r.db.Where("id = ?", id).Delete(data).Error
}

func (r *PropertyMysql) GetAll(req request.PropertyPaginationRequest) ([]model.Property, error) {
	var properties []model.Property
	offset := (req.Page - 1) * req.PageSize

	query := r.db.
		Table("properties").
		Joins("JOIN banks ON properties.bank_id = banks.id").
		Joins("JOIN selling_statuses ON properties.selling_status_id = selling_statuses.id").
		Joins("JOIN road_accesses ON properties.road_access_id = road_accesses.id").
		Joins("JOIN certificates ON properties.certificate_id = certificates.id").
		Preload("SellingStatus").
		Preload("Bank").
		Preload("PhotoHouse").
		Preload("PhotoCertificate").
		Preload("RoadAccess").
		Preload("Certificate").
		Limit(req.PageSize).
		Offset(offset)

	// Construct dynamic query based on available filters in the request
	if req.Query != "" {
		fields := []string{
			"title",
			"owner",
			"address",
			"building_area",
			"land_area",
			"latitude",
			"longitude",
			"property_tax_photo",
			"electricity_capacity",
			"water_source",
			"properties.description",
			"banks.name",
			"selling_statuses.name",
			"road_accesses.name",
			"certificates.name",
		}

		var conditions []string
		var values []interface{}

		for _, field := range fields {
			conditions = append(conditions, field+" LIKE ?")
			values = append(values, "%"+req.Query+"%")
		}

		if len(conditions) > 0 {
			query = query.Where(strings.Join(conditions, " OR "), values...)
		}
	}

	if err := query.
		Find(&properties).Error; err != nil {
		return nil, err
	}

	return properties, nil
}

func (r *PropertyMysql) GetById(id int64) (model.Property, error) {
	property := model.Property{}
	if err := r.db.
		Preload("SellingStatus").
		Preload("PhotoHouse").
		Preload("PhotoCertificate").
		Preload("RoadAccess").
		Preload("Certificate").
		Preload("Bank").
		First(&property, id).Error; err != nil {
		return property, err
	}

	return property, nil
}

func (r *PropertyMysql) GetBySellingStatus(id int64, page, pageSize int) (property []model.Property, err error) {
	offset := (page - 1) * pageSize
	if err := r.db.
		Where("selling_status_id = ?", id).
		Preload("SellingStatus").
		Preload("PhotoHouse").
		Preload("PhotoCertificate").
		Preload("RoadAccess").
		Preload("Bank").
		Preload("Certificate").
		Limit(pageSize).
		Offset(offset).
		Find(&property).Error; err != nil {
		return property, err
	}

	return property, nil
}

func (r *PropertyMysql) GetByLocation(latitude string, longitude string, radius string) ([]model.Property, error) {
	properties := []model.Property{}
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
	if err := r.db.Raw(sql, latitude, longitude, latitude, radius).
		Preload("Bank").
		Preload("SellingStatus").
		Preload("PhotoHouse").
		Preload("PhotoCertificate").
		Preload("RoadAccess").
		Preload("Certificate").
		Find(&properties).Error; err != nil {
		return properties, err
	}

	return properties, nil
}

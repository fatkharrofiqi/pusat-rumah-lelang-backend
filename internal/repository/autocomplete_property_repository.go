package repository

import (
	"pusat-rumah-lelang-backend/internal/model"
	"pusat-rumah-lelang-backend/internal/request"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IAutocompletePropertyRepository interface {
	GetComboProperty(db *gorm.DB, request *request.GetComboPropertyRequest) ([]*model.AutocompleteProperty, error)
}

type AutocompletePropertyRepository struct {
	Repository[model.Property]
	Log *logrus.Logger
}

func NewAutocompletePropertyRepository(log *logrus.Logger) IAutocompletePropertyRepository {
	return &AutocompletePropertyRepository{Log: log}
}

func (repo *AutocompletePropertyRepository) GetComboProperty(db *gorm.DB, request *request.GetComboPropertyRequest) ([]*model.AutocompleteProperty, error) {
	autoCompleteProperties := []*model.AutocompleteProperty{}
	offset := (request.Page - 1) * request.Size
	query := db.Table("properties").Select("title as label", "id").Limit(request.Size).Offset(offset)

	if request.Title != "" {
		query = query.Where("title LIKE ?", "%"+request.Title+"%")
	}

	if err := query.Scan(autoCompleteProperties).Error; err != nil {
		repo.Log.WithError(err).Error("failed to get autocomplete property repository")
		return autoCompleteProperties, err
	}

	return autoCompleteProperties, nil
}

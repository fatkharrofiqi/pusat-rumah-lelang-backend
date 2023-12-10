package autocomplete

import (
	"pusat-rumah-lelang-backend/models"
	"pusat-rumah-lelang-backend/requests"

	"gorm.io/gorm"
)

type IAutocompleteRepository interface {
	GetComboProperty(req *requests.AutocompletePaginationRequest) ([]models.AutocompleteProperty, error)
}

type AutocompleteRepository struct {
	db *gorm.DB
}

func NewAutocompleteRepository(db *gorm.DB) IAutocompleteRepository {
	return &AutocompleteRepository{db: db}
}

func (repo *AutocompleteRepository) GetComboProperty(req *requests.AutocompletePaginationRequest) ([]models.AutocompleteProperty, error) {
	var autoCompleteProperties []models.AutocompleteProperty
	offset := (req.Page - 1) * req.PageSize
	query := repo.db.Table("properties").Select("title as label", "id").Limit(req.PageSize).Offset(offset)

	if req.Title != "" {
		query = query.Where("title LIKE ?", "%"+req.Title+"%")
	}

	if err := query.Scan(&autoCompleteProperties).Error; err != nil {
		return autoCompleteProperties, err
	}

	return autoCompleteProperties, nil
}

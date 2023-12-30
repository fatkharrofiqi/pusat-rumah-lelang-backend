package autocomplete

import (
	"pusat-rumah-lelang-backend/internal/model"
	"pusat-rumah-lelang-backend/internal/request"

	"gorm.io/gorm"
)

type IAutocompleteRepository interface {
	GetComboProperty(req *request.AutocompletePaginationRequest) ([]model.AutocompleteProperty, error)
}

type AutocompleteRepository struct {
	db *gorm.DB
}

func NewAutocompleteRepository(db *gorm.DB) IAutocompleteRepository {
	return &AutocompleteRepository{db: db}
}

func (repo *AutocompleteRepository) GetComboProperty(req *request.AutocompletePaginationRequest) ([]model.AutocompleteProperty, error) {
	var autoCompleteProperties []model.AutocompleteProperty
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

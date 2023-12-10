package usecases

import (
	"pusat-rumah-lelang-backend/models"
	"pusat-rumah-lelang-backend/repositories/autocomplete"
	"pusat-rumah-lelang-backend/requests"
)

type IAutocompleteUsecase interface {
	GetComboProperty(req *requests.AutocompletePaginationRequest) ([]models.AutocompleteProperty, error)
}

type AutocompleteUsecase struct {
	repo autocomplete.IAutocompleteRepository
}

func NewAutocompleteUsecase(repo autocomplete.IAutocompleteRepository) IAutocompleteUsecase {
	return &AutocompleteUsecase{repo: repo}
}

func (u *AutocompleteUsecase) GetComboProperty(req *requests.AutocompletePaginationRequest) ([]models.AutocompleteProperty, error) {
	return u.repo.GetComboProperty(req)
}

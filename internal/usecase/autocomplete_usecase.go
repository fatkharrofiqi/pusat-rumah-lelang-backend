package usecase

import (
	"pusat-rumah-lelang-backend/internal/model"
	"pusat-rumah-lelang-backend/internal/repository"
	"pusat-rumah-lelang-backend/internal/request"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IAutocompletePropertyUsecase interface {
	GetComboProperty(c *gin.Context, request *request.GetComboPropertyRequest) ([]*model.AutocompleteProperty, error)
}

type AutocompletePropertyUsecase struct {
	AutocompletePropertyRepository repository.IAutocompletePropertyRepository
	DB                             *gorm.DB
}

func NewAutocompleteUsecase(db *gorm.DB, autoCompletePropertyRepository repository.IAutocompletePropertyRepository) IAutocompletePropertyUsecase {
	return &AutocompletePropertyUsecase{AutocompletePropertyRepository: autoCompletePropertyRepository, DB: db}
}

func (u *AutocompletePropertyUsecase) GetComboProperty(c *gin.Context, request *request.GetComboPropertyRequest) ([]*model.AutocompleteProperty, error) {
	return u.AutocompletePropertyRepository.GetComboProperty(u.DB.WithContext(c), request)
}

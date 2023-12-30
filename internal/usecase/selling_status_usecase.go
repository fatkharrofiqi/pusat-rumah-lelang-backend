package usecase

import (
	"pusat-rumah-lelang-backend/internal/model"
	"pusat-rumah-lelang-backend/internal/repository"
	"pusat-rumah-lelang-backend/internal/request"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ISellingStatusUsecase interface {
	Total(c *gin.Context) int64
	Create(c *gin.Context, request *request.CreateSellingStatusRequest) error
	Update(c *gin.Context, id int64, request *request.UpdateSellingStatusRequest) error
	Delete(c *gin.Context, id int64) error
	GetAll(c *gin.Context, request *request.GetAllSellingStatusRequest) ([]*model.SellingStatus, error)
	GetById(c *gin.Context, id int64) (*model.SellingStatus, error)
}

type SellingStatusUsecase struct {
	SellingStatusRepository repository.ISellingStatusRepository
	DB                      *gorm.DB
	Log                     *logrus.Logger
}

func NewSellingStatusUsecase(db *gorm.DB, sellingStatusRepository repository.ISellingStatusRepository, log *logrus.Logger) ISellingStatusUsecase {
	return &SellingStatusUsecase{DB: db, SellingStatusRepository: sellingStatusRepository, Log: log}
}

func (u *SellingStatusUsecase) Total(c *gin.Context) int64 {
	total := u.SellingStatusRepository.Total(u.DB.WithContext(c))
	return total
}

func (u *SellingStatusUsecase) Create(c *gin.Context, request *request.CreateSellingStatusRequest) error {
	return u.SellingStatusRepository.Create(u.DB.WithContext(c), request)
}

func (u *SellingStatusUsecase) Update(c *gin.Context, id int64, request *request.UpdateSellingStatusRequest) error {
	return u.SellingStatusRepository.Update(u.DB.WithContext(c), id, request)
}

func (u *SellingStatusUsecase) Delete(c *gin.Context, id int64) error {
	return u.SellingStatusRepository.Delete(u.DB.WithContext(c), id)
}

func (u *SellingStatusUsecase) GetAll(c *gin.Context, request *request.GetAllSellingStatusRequest) ([]*model.SellingStatus, error) {
	return u.SellingStatusRepository.GetAll(u.DB.WithContext(c), request)
}

func (u *SellingStatusUsecase) GetById(c *gin.Context, id int64) (*model.SellingStatus, error) {
	result, err := u.SellingStatusRepository.GetById(u.DB.WithContext(c), id)
	if err != nil {
		return result, err
	}

	return result, nil
}

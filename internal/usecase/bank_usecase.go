package usecase

import (
	"pusat-rumah-lelang-backend/internal/model"
	"pusat-rumah-lelang-backend/internal/repository"
	"pusat-rumah-lelang-backend/internal/request"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IBankUsecase interface {
	GetAll(ctx *gin.Context, request *request.GetAllBankRequest) (resp []*model.Bank, total int64, err error)
}

type BankUsecase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	BankRepository repository.IBankRepository
}

func NewBankUsecase(db *gorm.DB, bankRepository repository.IBankRepository, log *logrus.Logger) IBankUsecase {
	return &BankUsecase{
		DB:             db,
		Log:            log,
		BankRepository: bankRepository,
	}
}

func (b *BankUsecase) GetAll(ctx *gin.Context, request *request.GetAllBankRequest) (resp []*model.Bank, total int64, err error) {
	tx := b.DB.Begin()
	defer tx.Rollback()

	resp, total, err = b.BankRepository.GetAll(tx, request)
	if err != nil {
		b.Log.WithError(err).Error("error getting bank list")
		return resp, 0, err
	}

	if err := tx.Commit().Error; err != nil {
		b.Log.WithError(err).Error("error getting bank list")
		return resp, 0, err
	}

	return resp, total, err
}

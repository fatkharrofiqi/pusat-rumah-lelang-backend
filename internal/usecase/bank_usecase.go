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
	Create(ctx *gin.Context, request *request.CreateBankRequest) (err error)
	Update(ctx *gin.Context, request *request.UpdateBankRequest) (err error)
	Delete(ctx *gin.Context, request *request.DeleteBankRequest) (err error)
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

func (b *BankUsecase) Create(ctx *gin.Context, request *request.CreateBankRequest) (err error) {
	tx := b.DB.Begin()
	defer tx.Rollback()

	bank := &model.Bank{
		Name:        request.Name,
		Description: request.Description,
	}

	if err = b.BankRepository.Create(tx, bank); err != nil {
		b.Log.WithError(err).Error("Failed to create bank")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		b.Log.WithError(err).Error("Failed to commit")
		return err
	}

	return nil
}

func (b *BankUsecase) Update(ctx *gin.Context, request *request.UpdateBankRequest) (err error) {
	tx := b.DB.Begin()
	defer tx.Rollback()

	bank := &model.Bank{}
	if err = b.BankRepository.FindById(tx, bank, request.ID); err != nil {
		b.Log.WithError(err).Error("Failed to find bank")
		return err
	}

	bank.Name = request.Name
	bank.Description = request.Description

	if err = b.BankRepository.Update(tx, bank); err != nil {
		b.Log.WithError(err).Error("Failed to update bank")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		b.Log.WithError(err).Error("Failed to commit")
		return err
	}

	return nil
}

func (b *BankUsecase) Delete(ctx *gin.Context, request *request.DeleteBankRequest) (err error) {
	tx := b.DB.Begin()
	defer tx.Rollback()

	bank := &model.Bank{}
	if err = b.BankRepository.FindById(tx, bank, request.ID); err != nil {
		b.Log.WithError(err).Error("Failed to find bank")
		return err
	}

	if err = b.BankRepository.Delete(tx, bank); err != nil {
		b.Log.WithError(err).Error("Failed to delete bank")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		b.Log.WithError(err).Error("Failed to commit")
		return err
	}

	return nil
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

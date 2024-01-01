package usecase

import (
	"pusat-rumah-lelang-backend/internal/model"
	"pusat-rumah-lelang-backend/internal/repository"
	"pusat-rumah-lelang-backend/internal/request"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type IPropertyUsecase interface {
	GetTotalByCategory(c *gin.Context, category string) (result []*model.NameCount, err error)
	Create(c *gin.Context, request *request.CreatePropertyRequest) error
	Update(c *gin.Context, id int64, request *request.UpdatePropertyRequest) error
	Delete(c *gin.Context, id int64) error
	GetAll(c *gin.Context, request *request.GetAllPropertyRequest) (result []*model.Property, err error)
	GetById(c *gin.Context, id int64) (result *model.Property, err error)
	GetBySellingStatus(c *gin.Context, id int64, request *request.GetBySellingStatusRequest) (result []*model.Property, err error)
	GetByLocation(c *gin.Context, request *request.GetByLocationRequest) (result []*model.Property, err error)
}

type PropertyUsecase struct {
	PropertyRepository repository.IPropertyRepository
	DB                 *gorm.DB
	Log                *logrus.Logger
	Config             *viper.Viper
}

func NewPropertyUsecase(db *gorm.DB, repo repository.IPropertyRepository, log *logrus.Logger, config *viper.Viper) IPropertyUsecase {
	return &PropertyUsecase{
		DB:                 db,
		PropertyRepository: repo,
		Log:                log,
		Config:             config,
	}
}

func (p *PropertyUsecase) GetTotalByCategory(c *gin.Context, category string) (result []*model.NameCount, err error) {
	result, err = p.PropertyRepository.GetTotalByCategory(p.DB.WithContext(c), category)
	if err != nil {
		p.Log.WithError(err).Error("failed to retrieve category")
		return
	}
	return
}

func (p *PropertyUsecase) Create(c *gin.Context, request *request.CreatePropertyRequest) error {
	if err := p.PropertyRepository.Create(p.DB.WithContext(c), request); err != nil {
		p.Log.WithError(err).Error("failed to create category")
		return err
	}

	return nil
}

func (p *PropertyUsecase) Update(c *gin.Context, id int64, request *request.UpdatePropertyRequest) error {
	if err := p.PropertyRepository.Update(p.DB.WithContext(c), id, request); err != nil {
		p.Log.WithError(err).Error("failed to update category")
		return err
	}

	return nil
}

func (p *PropertyUsecase) Delete(c *gin.Context, id int64) error {
	if err := p.PropertyRepository.Delete(p.DB.WithContext(c), id); err != nil {
		p.Log.WithError(err).Error("failed to delete category")
		return err
	}

	return nil
}

// Function to process photo URLs
func (p *PropertyUsecase) processPhotoURLs(properties []*model.Property) {
	url := p.Config.GetString("MINIO_ENDPOINT_PROTOCOL") + "://" + p.Config.GetString("MINIO_ENDPOINT") + "/" + p.Config.GetString("MINIO_BUCKET") + "/"
	for i := range properties {
		for j := range properties[i].PhotoHouse {
			properties[i].PhotoHouse[j].PhotoUrl = url + strings.ReplaceAll(properties[i].PhotoHouse[j].PhotoUrl, " ", "%20")
		}
	}
}

func (p *PropertyUsecase) GetAll(c *gin.Context, req *request.GetAllPropertyRequest) (result []*model.Property, err error) {
	result, err = p.PropertyRepository.GetAll(p.DB.WithContext(c), req)
	if err != nil {
		p.Log.WithError(err).Error("failed to get all category")
		return
	}
	p.processPhotoURLs(result)
	return
}

func (p *PropertyUsecase) GetById(c *gin.Context, id int64) (result *model.Property, err error) {
	result, err = p.PropertyRepository.GetById(p.DB.WithContext(c), id)
	if err != nil {
		p.Log.WithError(err).Error("failed to get by id category")
		return
	}
	for index, photo := range result.PhotoHouse {
		result.PhotoHouse[index].PhotoUrl = "https://" + "MinioEndpoint" + "/" + "BucketName" + "/" + strings.ReplaceAll(photo.PhotoUrl, " ", "%20")
	}
	return
}

func (p *PropertyUsecase) GetBySellingStatus(c *gin.Context, id int64, request *request.GetBySellingStatusRequest) (results []*model.Property, err error) {
	results, err = p.PropertyRepository.GetBySellingStatus(p.DB.WithContext(c), id, request)
	if err != nil {
		p.Log.WithError(err).Error("failed to get selling by status")
		return
	}
	p.processPhotoURLs(results)
	return
}

func (p *PropertyUsecase) GetByLocation(c *gin.Context, request *request.GetByLocationRequest) (result []*model.Property, err error) {
	result, err = p.PropertyRepository.GetByLocation(p.DB.WithContext(c), request)
	if err != nil {
		p.Log.WithError(err).Error("failed to get by location")
		return
	}
	p.processPhotoURLs(result)
	return
}

package usecases

import (
	"encoding/json"
	"os"
	"pusat-rumah-lelang-backend/common/interfaces"
	"pusat-rumah-lelang-backend/models"
)

type IBannerUsecase interface {
	interfaces.IGetAllGeneric[models.Banner]
}

type BannerUsecase struct{}

func NewBannerUsecase() IBannerUsecase {
	return &BannerUsecase{}
}

func (b *BannerUsecase) GetAll(page, pageSize int) ([]models.Banner, error) {
	file, err := os.Open("data/banner-url.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data map[string][]string
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return nil, err
	}

	urls, ok := data["url"]
	if !ok {
		return nil, err
	}

	var banners []models.Banner
	for _, url := range urls {
		banners = append(banners, models.Banner{
			Url: url,
		})
	}

	return banners, nil
}

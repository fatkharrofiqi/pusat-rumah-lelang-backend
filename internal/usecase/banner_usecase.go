package usecase

import (
	"encoding/json"
	"os"
	"pusat-rumah-lelang-backend/internal/common/interfaces"
	"pusat-rumah-lelang-backend/internal/model"
)

type IBannerUsecase interface {
	interfaces.IGetAllGeneric[model.Banner]
}

type BannerUsecase struct{}

func NewBannerUsecase() IBannerUsecase {
	return &BannerUsecase{}
}

func (b *BannerUsecase) GetAll(page, pageSize int) ([]model.Banner, error) {
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

	var banners []model.Banner
	for _, url := range urls {
		banners = append(banners, model.Banner{
			Url: url,
		})
	}

	return banners, nil
}

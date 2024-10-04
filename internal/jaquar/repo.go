package jaquar

import "github.com/MetsysEht/Tiles-Invoice-BE/internal/boot"

type repo struct{}

func NewRepo() IRepository {
	return &repo{}
}

func (r *repo) GetProductDetails(series string, colorCode string, codeNumber string) (*Product, error) {
	product := Product{}
	q := boot.DB.Where("series = ? AND color_code = ? AND code_number = ?", series, colorCode, codeNumber).First(&product)
	if q.Error != nil {
		return nil, q.Error
	}
	return &product, nil
}

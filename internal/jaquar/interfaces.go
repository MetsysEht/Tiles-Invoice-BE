package jaquar

type IRepository interface {
	GetProductDetails(series string, colorCode string, codeNumber string) (*Product, error)
}

type IManager interface {
	GetProductDetails(series string, colorCode string, codeNumber string) (*Product, error)
}

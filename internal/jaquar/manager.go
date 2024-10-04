package jaquar

type manager struct {
	repo IRepository
}

func NewManager(repo IRepository) IManager {
	return &manager{repo}
}

func (m manager) GetProductDetails(series string, colorCode string, codeNumber string) (*Product, error) {
	return m.repo.GetProductDetails(series, colorCode, codeNumber)
}

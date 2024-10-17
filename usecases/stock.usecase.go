package usecases

import (
	"github.com/muhammadsarimin/indocyber-api/config"
	"github.com/muhammadsarimin/indocyber-api/models"
	"github.com/muhammadsarimin/indocyber-api/repositories"
)

type StockUsecase interface {
	GetStocks() ([]models.Stock, error)
	GetStock(id int) (*models.Stock, error)
	CreateStock(stock *models.Stock) (*models.Stock, error)
	UpdateStock(stock *models.Stock) (*models.Stock, error)
	DeleteStock(id int) error
	StockExist(id int) bool
}

type stockUsecase struct {
	repo repositories.StockRepo
	log  *config.CustomLog
}

func NewStockUsecase(repo repositories.StockRepo, log *config.CustomLog) StockUsecase {
	return &stockUsecase{repo, log}
}

func (u *stockUsecase) GetStocks() ([]models.Stock, error) {
	return u.repo.GetStocks()
}

func (u *stockUsecase) GetStock(id int) (*models.Stock, error) {
	return u.repo.GetStock(id)
}

func (u *stockUsecase) CreateStock(stock *models.Stock) (*models.Stock, error) {
	return u.repo.CreateStock(stock)
}

func (u *stockUsecase) UpdateStock(stock *models.Stock) (*models.Stock, error) {
	return u.repo.UpdateStock(stock)
}

func (u *stockUsecase) DeleteStock(id int) error {
	return u.repo.DeleteStock(id)
}

func (u *stockUsecase) StockExist(id int) bool {
	stock, err := u.repo.GetStock(id)
	if err != nil {
		return false
	}

	return stock.ID > 0
}

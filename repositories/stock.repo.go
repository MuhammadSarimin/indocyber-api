package repositories

import (
	"github.com/muhammadsarimin/indocyber-api/config"
	"github.com/muhammadsarimin/indocyber-api/models"
	"gorm.io/gorm"
)

type StockRepo interface {
	GetStocks() ([]models.Stock, error)
	GetStock(id int) (*models.Stock, error)
	CreateStock(stock *models.Stock) (*models.Stock, error)
	UpdateStock(stock *models.Stock) (*models.Stock, error)
	DeleteStock(id int) error
}

type stockRepo struct {
	db  *gorm.DB
	log *config.CustomLog
}

func NewStockRepo(db *gorm.DB, log *config.CustomLog) StockRepo {
	return &stockRepo{db, log}
}

func (r *stockRepo) GetStocks() ([]models.Stock, error) {

	var stocks []models.Stock
	err := r.db.Find(&stocks).Error

	if err != nil {
		r.log.Error(err, "GetStocks", nil, nil)
	}

	return stocks, err

}

func (r *stockRepo) GetStock(id int) (*models.Stock, error) {

	var stock models.Stock
	err := r.db.First(&stock, id).Error

	if err != nil {
		r.log.Error(err, "GetStock", nil, nil)
	}

	return &stock, err

}

func (r *stockRepo) CreateStock(stock *models.Stock) (*models.Stock, error) {
	err := r.db.Create(&stock).Error
	if err != nil {
		r.log.Error(err, "CreateStock", nil, nil)
	}

	return stock, err

}

func (r *stockRepo) UpdateStock(stock *models.Stock) (*models.Stock, error) {

	err := r.db.Save(&stock).Error
	if err != nil {
		r.log.Error(err, "UpdateStock", nil, nil)
	}

	return stock, err

}

func (r *stockRepo) DeleteStock(id int) error {

	var stock models.Stock
	err := r.db.Delete(&stock).Where("id = ?", id).Error
	if err != nil {
		r.log.Error(err, "DeleteStock", nil, nil)
	}

	return err

}

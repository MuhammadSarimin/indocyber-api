package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/muhammadsarimin/indocyber-api/models"
	"github.com/muhammadsarimin/indocyber-api/usecases"
	"github.com/stretchr/testify/assert"
)

var app *gin.Engine

func TestMain(m *testing.M) {
	app = gin.Default()
	api := app.Group("/api/v1")

	repo := &mockStockRepo{}
	usecase := usecases.NewStockUsecase(repo, nil)
	NewStockHandler(api, usecase, nil)

}

func TestGetStocks(t *testing.T) {
	t.Run("Get all stock success", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/api/v1/stocks", nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		app.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

	})
}

type mockStockRepo struct{}

func (r *mockStockRepo) GetStocks() ([]models.Stock, error) {
	return nil, nil
}

func (r *mockStockRepo) GetStock(id int) (*models.Stock, error) {
	return nil, nil
}

func (r *mockStockRepo) CreateStock(stock *models.Stock) (*models.Stock, error) {
	return nil, nil
}

func (r *mockStockRepo) UpdateStock(stock *models.Stock) (*models.Stock, error) {
	return nil, nil
}

func (r *mockStockRepo) DeleteStock(id int) error {
	return nil
}

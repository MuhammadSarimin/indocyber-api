package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadsarimin/indocyber-api/config"
	"github.com/muhammadsarimin/indocyber-api/handlers"
	"github.com/muhammadsarimin/indocyber-api/middleware"
	"github.com/muhammadsarimin/indocyber-api/repositories"
	"github.com/muhammadsarimin/indocyber-api/usecases"
)

func main() {

	config.Init()
	app := gin.Default()

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"response_code": "404",
			"response_msg":  "Not Found",
		})
	})

	api := app.Group("/api/v1")
	api.Use(middleware.BasicAuth())

	logger := config.NewCustomLog()

	handlers.NewStockHandler(
		api,
		usecases.NewStockUsecase(
			repositories.NewStockRepo(
				config.NewDB(
					config.Env.DB,
				),
				logger,
			),
			logger,
		),
		logger,
	)

	app.Run(config.Env.Address())
}

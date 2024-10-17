package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhammadsarimin/indocyber-api/config"
	"github.com/muhammadsarimin/indocyber-api/helper"
	"github.com/muhammadsarimin/indocyber-api/helper/response"
	"github.com/muhammadsarimin/indocyber-api/models"
	"github.com/muhammadsarimin/indocyber-api/models/cerr"
	"github.com/muhammadsarimin/indocyber-api/usecases"
)

type stockHandler struct {
	usecase usecases.StockUsecase
	log     *config.CustomLog
}

func NewStockHandler(r *gin.RouterGroup, usecase usecases.StockUsecase, log *config.CustomLog) {
	handler := &stockHandler{usecase, log}
	r.GET("/stocks", handler.GetStocks)
	r.GET("/stocks/:id", handler.GetStock)
	r.POST("/stocks", handler.CreateStock)
	r.PUT("/stocks/:id", handler.UpdateStock)
	r.DELETE("/stocks/:id", handler.DeleteStock)
}

func (h *stockHandler) GetStocks(c *gin.Context) {
	stocks, err := h.usecase.GetStocks()
	if err != nil {
		response.Error(c, h.log, err)
		return
	}

	response.Success(c, h.log, stocks)
}

func (h *stockHandler) GetStock(c *gin.Context) {
	param := c.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		h.log.Error(err, "GetStock", nil, nil)
		response.Error(c, h.log, cerr.GetError("003", ""))
		return
	}

	stock, err := h.usecase.GetStock(id)
	if err != nil {
		h.log.Error(err, "GetStock", nil, nil)
		response.Error(c, h.log, err)
		return
	}

	response.Success(c, h.log, stock)
}

func (h *stockHandler) CreateStock(c *gin.Context) {

	payload := models.Stock{}
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		h.log.Error(err, "CreateStock", nil, nil)
		response.Error(c, h.log, err)
		return
	}

	if err := helper.Validate(payload); err != nil {
		h.log.Error(err, "CreateStock", nil, nil)
		response.Error(c, h.log, err)
		return
	}

	payload.CreatedBy = h.user(c)
	c.Set("payload", payload)

	stock, err := h.usecase.CreateStock(&payload)
	if err != nil {
		h.log.Error(err, "CreateStock", payload, nil)
		response.Error(c, h.log, err)
		return
	}

	h.log.Info("CreateStock", payload, nil)
	response.Success(c, h.log, stock)
}

func (h *stockHandler) UpdateStock(c *gin.Context) {

	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		h.log.Error(err, "UpdateStock", nil, nil)
		response.Error(c, h.log, cerr.GetError("003", ""))
		return
	}

	if !h.usecase.StockExist(id) {
		h.log.Error(errors.New("data not found"), "UpdateStock", nil, nil)
		response.Error(c, h.log, cerr.GetError("004", ""))
		return
	}

	payload := models.Stock{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		h.log.Error(err, "UpdateStock", nil, nil)
		response.Error(c, h.log, err)
		return
	}

	payload.ID = uint(id)
	payload.UpdatedBy = h.user(c)
	c.Set("payload", payload)

	stock, err := h.usecase.UpdateStock(&payload)
	if err != nil {
		h.log.Error(err, "UpdateStock", payload, nil)
		response.Error(c, h.log, err)
		return
	}

	h.log.Info("UpdateStock", payload, nil)
	response.Success(c, h.log, stock)
}

func (h *stockHandler) DeleteStock(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		h.log.Error(err, "DeleteStock", nil, nil)
		response.Error(c, h.log, cerr.GetError("003", ""))
		return
	}

	if !h.usecase.StockExist(id) {
		h.log.Error(errors.New("data not found"), "DeleteStock", nil, nil)
		response.Error(c, h.log, cerr.GetError("004", ""))
		return
	}

	err = h.usecase.DeleteStock(id)
	if err != nil {
		h.log.Error(err, "DeleteStock", nil, nil)
		response.Error(c, h.log, err)
		return
	}

	response.Success(c, h.log, nil)
}

func (h *stockHandler) user(c *gin.Context) string {
	user, _, _ := c.Request.BasicAuth()
	return user
}

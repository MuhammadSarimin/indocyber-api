package models

import (
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	ID             uint           `json:"id"              gorm:"primarykey"`
	NamaBarang     string         `json:"nama_barang"     validate:"required"`
	Quantity       int            `json:"quantity"        validate:"required"`
	Seri           string         `json:"seri"            validate:"required"`
	AdditionalInfo JSONMap        `json:"additional_info" validate:"required" gorm:"type:json"`
	CreatedAt      time.Time      `json:"created_at"`
	CreatedBy      string         `json:"created_by"`
	UpdatedAt      time.Time      `json:"updated_at"`
	UpdatedBy      string         `json:"updated_by"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}

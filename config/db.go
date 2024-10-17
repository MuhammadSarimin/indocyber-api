package config

import (
	"log"

	"github.com/muhammadsarimin/indocyber-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(con models.DBConfig) *gorm.DB {

	db, err := gorm.Open(postgres.Open(con.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug()

	if con.AutoMigrate {
		db.AutoMigrate(&models.Stock{})
	}

	return db

}

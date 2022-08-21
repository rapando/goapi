package db

import (
	"log"

	"github.com/rapando/goapi/pkg/common/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(uri string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})
	return db
}

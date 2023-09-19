package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Deklarasi Database dengan nama DB
var DB *gorm.DB

func ConnectDatabase() {
	// dsn (Data Source Name) [reference: https://github.com/go-sql-driver/mysql#dsn-data-source-name]
	dsn := "root:@tcp(localhost:3306)/db_restfulapi_gin"
	database, error := gorm.Open(mysql.Open(dsn))
	if error != nil {
		panic(error)
	}

	// Migrate ke struct models/product.go
	database.AutoMigrate(&Product{})

	// Masukkan database ke DB (gorm.DB)
	DB = database
}

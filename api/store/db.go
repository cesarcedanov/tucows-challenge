package store

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"tucows-challenge/api/model"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err = db.AutoMigrate(&model.Order{}); err != nil {
		panic(err)
	}
	db.Model(&model.Order{}).Create(InitOrders)
	log.Println("Successfully connected to DB")
	return db
}

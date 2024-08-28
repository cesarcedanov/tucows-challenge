package store

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"tucows-challenge/api/model"
)

func InitDB() *gorm.DB {
	//dsn := "host=0.0.0.0 user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://user:password@localhost:5432/mydb?sslmode=disable"
		log.Println("DATABASE_URL environment variable not set", "Will use the following one:", dsn)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err = db.AutoMigrate(&model.Order{}); err != nil {
		panic(err)
	}
	db.Model(&model.Order{}).Create(model.InitOrders)
	log.Println("Successfully connected to DB")
	return db
}

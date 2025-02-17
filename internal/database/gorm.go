package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ShortUrl struct {
	ID uint `gorm:"primaryKey"`
	Short string `gorm:"index:,type:hash"`
	Original string `gorm:"index:,type:hash"`
}

var DB *gorm.DB

func Connect() error {
	dsn := "host=192.168.68.104 user=admin password=admin dbname=shorturl port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	log.Print("Применение миграций")

	err = db.AutoMigrate(ShortUrl{})

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Миграции применены")
	DB = db
	return nil
}

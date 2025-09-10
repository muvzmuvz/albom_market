package db

import (
	"fmt"
	"gin/sturct"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dsn := "host=localhost user=ginuser password=ginpass dbname=gindb port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("ошибка подключения: %w", err)
	}

	err = DB.AutoMigrate(&sturct.Album{}, &sturct.User{})
	if err != nil {
		return fmt.Errorf("ошибка миграции: %w", err)
	}

	return nil
}

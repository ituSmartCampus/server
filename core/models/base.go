package models

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {

	DB, err := gorm.Open(sqlite.Open("db/test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return DB
}

var DB = ConnectDB()

type Base struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DBConn() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

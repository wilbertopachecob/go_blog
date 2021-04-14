package db

import (
	"wilbertopachecob/go_blog/cmd/web/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

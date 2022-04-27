package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func dbInit(
	user string,
	password string,
	host string,
	port int,
	database string,
) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Asia%%2FTokyo&charset=utf8mb4",
		user,
		password,
		host,
		port,
		database,
	)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.AutoMigrate(&User{}, &Memo{})
	if err != nil {
		return fmt.Errorf("failed to auto migrate: %w", err)
	}

	return nil
}

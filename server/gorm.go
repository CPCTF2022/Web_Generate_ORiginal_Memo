package main

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db        *gorm.DB
	errNoUser = errors.New("no user")
	errNoMemo = errors.New("no memo")
)

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

func createUser(ctx context.Context, user *User) error {
	return db.
		WithContext(ctx).
		Create(user).Error
}

func getUserByName(ctx context.Context, userName string) (*User, error) {
	var user User
	err := db.
		WithContext(ctx).
		First(&user, User{Name: userName}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errNoUser
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func createMemo(ctx context.Context, memo *Memo) error {
	return db.
		WithContext(ctx).
		Create(memo).Error
}

func getMemos(ctx context.Context, userID string) ([]Memo, error) {
	var memos []Memo
	err := db.
		WithContext(ctx).
		Find(&memos, Memo{UserID: userID}).Error
	if err != nil {
		return nil, err
	}

	return memos, nil
}

func getMemo(ctx context.Context, memoID string) (*Memo, error) {
	var memo Memo
	err := db.
		WithContext(ctx).
		First(&memo, memoID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errNoMemo
	}
	if err != nil {
		return nil, err
	}

	return &memo, nil
}

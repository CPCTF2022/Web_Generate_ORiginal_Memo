package main

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db        *gorm.DB
	errNoUser = errors.New("no user")
	errNoMemo = errors.New("no memo")
)

func dbInit(
	flag string,
	adminPassword string,
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to generate hashed password: %w", err)
	}
	admin := User{
		Name:           "admin",
		HashedPassword: string(hashedPassword),
	}
	err = db.
		Session(&gorm.Session{}).
		FirstOrCreate(&admin, "`name` = \"admin\"").Error
	if err != nil {
		return fmt.Errorf("failed to create admin user: %w", err)
	}

	err = db.
		Session(&gorm.Session{}).
		Where("user_id = ?", admin.ID).
		FirstOrCreate(&Memo{
			UserID:  admin.ID,
			Content: flag,
		}).Error
	if err != nil {
		return fmt.Errorf("failed to create admin memo: %w", err)
	}

	return nil
}

func createUser(ctx context.Context, user *User) error {
	return db.
		WithContext(ctx).
		Create(user).Error
}

func getUserByName(ctx context.Context, userName string) (*User, string, error) {
	var user User
	db := db.
		WithContext(ctx).
		Where("name = ?", userName)
	query := db.ToSQL(func(db *gorm.DB) *gorm.DB {
		return db.First(&user)
	})
	err := db.First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, query, errNoUser
	}
	if err != nil {
		return nil, query, err
	}

	return &user, query, nil
}

func createMemo(ctx context.Context, memo *Memo) error {
	return db.
		WithContext(ctx).
		Create(memo).Error
}

func getMemos(ctx context.Context, userID int) ([]Memo, string, error) {
	var memos []Memo
	db := db.
		WithContext(ctx).
		Where("user_id = ?", userID)
	query := db.ToSQL(func(db *gorm.DB) *gorm.DB {
		return db.Find(&memos)
	})
	err := db.Find(&memos).Error
	if err != nil {
		return nil, query, err
	}

	return memos, query, nil
}

func getMemo(ctx context.Context, memoID string, userID int) (*Memo, string, error) {
	var memo Memo
	db := db.
		WithContext(ctx).
		Where(memoID).
		Where("user_id = ?", userID)
	query := db.ToSQL(func(db *gorm.DB) *gorm.DB {
		return db.First(&memo)
	})
	err := db.First(&memo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, query, errNoMemo
	}
	if err != nil {
		return nil, query, err
	}

	return &memo, query, nil
}

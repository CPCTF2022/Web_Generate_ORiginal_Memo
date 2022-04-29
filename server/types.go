package main

import (
	"time"
)

type User struct {
	ID             int    `gorm:"autoIncrement;not null;primaryKey" json:"id"`
	Name           string `gorm:"type:varchar(50);not null;size:50;unique" json:"name"`
	Password       string `gorm:"-" json:"password"`
	HashedPassword string `gorm:"type:char(60);not null;size:60" json:"-"`
}

type Memo struct {
	ID        int       `gorm:"autoIncrement;not null;primaryKey" json:"id"`
	UserID    int       `gorm:"not null;size:36" json:"user_id"`
	Content   string    `gorm:"type:varchar(255);not null;size:255" json:"content"`
	CreatedAt time.Time `gorm:"type:timestamp;not null" json:"createdAt"`
}

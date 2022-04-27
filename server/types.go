package main

import (
	"time"
)

type User struct {
	ID             string `gorm:"type:char(36);not null;primaryKey;size:36" json:"id"`
	Name           string `gorm:"type:varchar(50);not null;size:50;unique" json:"name"`
	Password       string `gorm:"-" json:"password"`
	HashedPassword string `gorm:"type:char(60);not null;size:60" json:"-"`
}

type Memo struct {
	ID        string    `gorm:"type:char(36);not null;primaryKey;size:36" json:"id"`
	UserID    string    `gorm:"type:char(36);not null;size:36" json:"user_id"`
	Content   string    `gorm:"type:varchar(255);not null;size:255" json:"content"`
	CreatedAt time.Time `gorm:"type:timestamp;not null" json:"created_at"`
}

package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Email    string `gorm:"uniqueIndex;type:varchar(128) not null"`
	Password string `gorm:"type:varchar(64) not null"`
}


func (u *User) TableName() string {
	return "user"
}
package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string `json:"id" gorm:"primaryKey;type:char(14)"`
	Username  string `gorm:"type:varchar(100);unique"`
	Password  string `gorm:"size:255"`
	Role      string `gorm:"type:varchar(10)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

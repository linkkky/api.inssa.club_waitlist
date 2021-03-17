package models

import (
	"time"

	"gorm.io/gorm"
)

// Interest is a model, which is the entity of waitlist
type Interest struct {
	ID              uint64         `gorm:"primary_key"`
	ClubhouseUserID uint64         `gorm:"column:clubhouse_user_id;unique;"`
	Email           string         `gorm:"column:email;unique;not null"`
	CreatedAt       time.Time      `gorm:"column:created_at;autoCreateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"column:delete_at"`
}

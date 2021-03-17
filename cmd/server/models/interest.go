package models

import (
	"gorm.io/gorm"
)

// Interest is a model, which is the entity of waitlist
type Interest struct {
	gorm.Model
	ID              uint64 `gorm:"primary_key"`
	ClubhouseUserID uint64 `gorm:"column:clubhouse_user_id;unique;"`
	Email           string `gorm:"column:email;unique;not null" valid:"email"`
}

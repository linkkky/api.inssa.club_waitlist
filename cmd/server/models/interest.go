package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Interest is a model, which is the entity of waitlist
type Interest struct {
	gorm.Model
	ID              uint64 `gorm:"primary_key" example:"3"`
	ClubhouseUserID uint64 `gorm:"column:clubhouse_user_id;unique;" example:"123456"`
	Email           string `gorm:"column:email;unique;not null" valid:"email" example:"example@example.com"`
}

func (interest *Interest) Create(db *gorm.DB) error {
	isValid, err := govalidator.ValidateStruct(interest)
	if isValid {
		db.Create(&interest)
		return nil
	}
	return err
}

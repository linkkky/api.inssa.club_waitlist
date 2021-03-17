package models

import (
	"gopkg.in/guregu/null.v4"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Interest is a model, which is the entity of waitlist
type Interest struct {
	gorm.Model
	ClubhouseUserID null.Int `json:"clubhouse_user_id" gorm:"column:clubhouse_user_id;unique;" example:"123456"`
	Email           string   `json:"email" gorm:"column:email;unique;not null" valid:"email, required" example:"example@example.com"`
}

func (interest *Interest) Create(db *gorm.DB) error {
	isValid, err := govalidator.ValidateStruct(interest)
	if isValid {
		db.Create(&interest)
		return nil
	}
	return err
}

package forms

import "gopkg.in/guregu/null.v4"

// AddInterest is a struct for adding interest request
type AddInterest struct {
	ClubhouseUserID null.Int `json:"clubhouse_user_id" valid:"optional"`
	Email           string   `json:"email" valid:"email, required" binding:"required"`
}

// DeleteInterest is a struct for deleting interest request
type DeleteInterest struct {
	Email string `json:"email" valid:"email, required" binding:"required"`
}

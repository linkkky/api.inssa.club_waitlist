package forms

import "gopkg.in/guregu/null.v4"

// AddInterestRequest is a struct for adding interest request
type AddInterestRequest struct {
	ClubhouseUserID null.Int `json:"clubhouse_user_id" valid:"optional"`
	Email           string   `json:"email" valid:"email, required" binding:"required"`
}

// AddInterestResponse is a struct for adding interest response
type AddInterestResponse struct {
	ID              int64
	ClubhouseUserID null.Int `json:"clubhouse_user_id" valid:"optional"`
	Email           string   `json:"email" valid:"email, required" binding:"required"`
}

// DeleteInterest is a struct for deleting interest request
type DeleteInterest struct {
	Email string `json:"email" valid:"email, required" binding:"required"`
}

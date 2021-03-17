package forms

// AddInterest is a struct for adding interest request
type AddInterest struct {
	UserID string `json:"user_id" valid:"optional"`
	Email  string `json:"email" valid:"email, required" binding:"required"`
}

// DeleteInterest is a struct for deleting interest request
type DeleteInterest struct {
	Email string `json:"email" valid:"email, required" binding:"required"`
}

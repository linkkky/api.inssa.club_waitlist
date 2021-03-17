package forms

// AddInterest is a struct for adding interest request
type AddInterest struct {
	UserID string `json:"user_id" validate:"optional"`
	Email  string `json:"email" validate:"email, required"`
}

// DeleteInterest is a struct for deleting interest request
type DeleteInterest struct {
	Email string `json:"email" validate:"email, required"`
}

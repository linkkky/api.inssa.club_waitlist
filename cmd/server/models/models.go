package models

// GetModels returns the registered models
func GetModels() []interface{} {
	models := []interface{}{&Interest{}}
	return models
}

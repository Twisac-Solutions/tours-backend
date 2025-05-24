package models

// ErrorResponse represents a standard error response.
// swagger:model
type ErrorResponse struct {
	Error string `json:"error"`
}

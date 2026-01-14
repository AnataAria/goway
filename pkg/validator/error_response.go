package validator

import "encoding/json"

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Errors  []FieldError `json:"errors,omitempty"`
}

type ValidationError struct {
	*ErrorResponse
}

func (e *ValidationError) Error() string {
	b, _ := json.Marshal(e.ErrorResponse)
	return string(b)
}

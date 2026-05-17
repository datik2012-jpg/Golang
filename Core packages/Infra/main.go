package main

import (
	"fmt"
)

type APIError struct {
	Message string
	Code    int
	Details map[string]interface{}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("APIError %d: %s | details=%v", e.Code, e.Message, e.Details)
}

func NewAPIError(code int, message string, details map[string]interface{}) error {
	return &APIError{
		Message: message,
		Code:    code,
		Details: details,
	}
}

func main() {
	err := NewAPIError(400, "Bad request", map[string]interface{}{
		"field": "username",
		"error": "cannot be empty",
	})
	fmt.Println(err.Error())
}

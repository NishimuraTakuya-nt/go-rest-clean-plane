package apperrors

import (
	"fmt"
	"strings"
)

// FieldError represents a validation error for a specific field
type FieldError struct {
	Field   string
	Message string
}

// ValidationErrors is a collection of FieldErrors
type ValidationErrors []FieldError

func (ve *ValidationErrors) Error() string {
	var errMessages []string
	for _, fe := range *ve {
		errMessages = append(errMessages, fmt.Sprintf("%s: %s", fe.Field, fe.Message))
	}
	return strings.Join(errMessages, "; ")
}

// AddError adds a new FieldError to ValidationErrors
func (ve *ValidationErrors) AddError(field, message string) {
	*ve = append(*ve, FieldError{Field: field, Message: message})
}

// NewValidationErrors creates a new ValidationErrors instance
func NewValidationErrors() *ValidationErrors {
	return &ValidationErrors{}
}

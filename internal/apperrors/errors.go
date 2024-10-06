package apperrors

import (
	"net/http"
)

// AppError はアプリケーション固有のエラーを表します。
type AppError struct {
	Type       ErrorType
	RawError   error
	StatusCode int
	Message    string
}

func (e *AppError) Error() string {
	return e.Message
}

// ErrorType はエラーの種類を表す列挙型です。
type ErrorType string

const (
	ErrorTypeValidation         ErrorType = "VALIDATION_ERROR"
	ErrorTypeNotFound           ErrorType = "NOT_FOUND"
	ErrorTypeUnauthorized       ErrorType = "UNAUTHORIZED"
	ErrorTypeForbidden          ErrorType = "FORBIDDEN"
	ErrorTypeInternal           ErrorType = "INTERNAL_ERROR"
	ErrorTypeBadRequest         ErrorType = "BAD_REQUEST"
	ErrorTypeConflict           ErrorType = "CONFLICT"
	ErrorTypeRateLimit          ErrorType = "RATE_LIMIT"
	ErrorTypeServiceUnavailable ErrorType = "SERVICE_UNAVAILABLE"
	ErrorTypeExternalService    ErrorType = "EXTERNAL_SERVICE_ERROR"
	ErrorTypeDataAccess         ErrorType = "DATA_ACCESS_ERROR"
	ErrorTypeBusinessLogic      ErrorType = "BUSINESS_LOGIC_ERROR"
)

// NewAppError creates a new AppError
func NewAppError(errType ErrorType, rawErr error, statusCode int, message string) *AppError {
	return &AppError{
		Type:       errType,
		RawError:   rawErr,
		StatusCode: statusCode,
		Message:    message,
	}
}

// Predefined error creators
func NewValidationError(message string, rawErr error) *AppError {
	return NewAppError(ErrorTypeValidation, rawErr, http.StatusBadRequest, message)
}

func NewNotFoundError(message string, rawErr error) *AppError {
	return NewAppError(ErrorTypeNotFound, rawErr, http.StatusNotFound, message)
}

func NewUnauthorizedError(message string, rawErr error) *AppError {
	return NewAppError(ErrorTypeUnauthorized, rawErr, http.StatusUnauthorized, message)
}

func NewForbiddenError(message string, rawErr error) *AppError {
	return NewAppError(ErrorTypeForbidden, rawErr, http.StatusForbidden, message)
}

func NewInternalError(message string, rawErr error) *AppError {
	return NewAppError(ErrorTypeInternal, rawErr, http.StatusInternalServerError, message)
}

func NewBadRequestError(message string, rawErr error) *AppError {
	return NewAppError(ErrorTypeBadRequest, rawErr, http.StatusBadRequest, message)
}

func NewConflictError(message string, rawErr error) *AppError {
	return NewAppError(ErrorTypeConflict, rawErr, http.StatusConflict, message)
}

func NewRateLimitError(message string, rawErr error) *AppError {
	return NewAppError(ErrorTypeRateLimit, rawErr, http.StatusTooManyRequests, message)
}

func NewServiceUnavailableError(message string, rawErr error) *AppError {
	return NewAppError(ErrorTypeServiceUnavailable, rawErr, http.StatusServiceUnavailable, message)
}

func NewExternalServiceError(message string, rawErr error) *AppError {
	return NewAppError(ErrorTypeExternalService, rawErr, http.StatusBadGateway, message)
}

func NewDataAccessError(message string, rawErr error) *AppError {
	return NewAppError(ErrorTypeDataAccess, rawErr, http.StatusInternalServerError, message)
}

func NewBusinessLogicError(message string, rawErr error) *AppError {
	return NewAppError(ErrorTypeBusinessLogic, rawErr, http.StatusUnprocessableEntity, message)
}

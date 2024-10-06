package apperrors

type ErrorWriter interface {
	WriteError(err error)
}

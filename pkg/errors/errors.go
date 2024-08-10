package errors

// Type used to add context to errors without using reflection
type ErrorId string

// Type used to infer the http status code to return and handle
// specific errors without using errors.Is (expensive)
type ErrorType int

const (
	AddMessageError    ErrorId = "AddMessageError"
	UnmarshalJSONError ErrorId = "UnmarshalJSONError"
	BytesReadingError  ErrorId = "BytesReadingError"
)

const (
	InternalError ErrorType = iota
	NotFoundError ErrorType = iota
)

package utils

// A custom response to handle response for methods within the business logic
type CustomResponse struct {
	Err     error
	Code    uint
	Message string
}

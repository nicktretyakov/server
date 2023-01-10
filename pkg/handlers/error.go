package handlers

type APIError struct {
	Message string
}

func NewAPIError(err error) APIError {
	return APIError{Message: err.Error()}
}

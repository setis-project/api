package models

type ApiError struct {
	Error string `json:"error"`
}

type ApiRequestError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ApiRequestErrors struct {
	Errors []ApiRequestError `json:"errors"`
}

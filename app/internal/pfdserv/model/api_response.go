package model

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ApiError   `json:"error,omitempty"`
}

type ApiError struct {
	ErrorTitle   string `json:"error_title"`
	ErrorMessage string `json:"error_message"`
}

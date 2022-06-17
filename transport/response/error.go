package response

type ResponseError struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
}

func (e *ResponseError) ErrorResponse(code int, msg string) *ResponseError {
	return &ResponseError{
		Status:  code,
		Message: msg,
	}
}

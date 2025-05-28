package errors

type HTTPExceptionResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Cause      string `json:"cause,omitempty"`
}

func NewHTTPDefaultExceptionResponse(statusCode int, message string) HTTPExceptionResponse {
	return HTTPExceptionResponse{
		StatusCode: statusCode,
		Message:    message,
	}
}

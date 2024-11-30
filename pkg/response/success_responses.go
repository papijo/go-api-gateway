package response

func SuccessResponse(message string, data interface{}) *BaseResponse {
	return &BaseResponse{
		Error:   false,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(message string, err interface{}) *BaseResponse {
	return &BaseResponse{
		Error:   true,
		Message: message,
		Data:    err,
	}
}

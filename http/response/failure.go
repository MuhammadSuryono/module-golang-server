package response

import models "github.com/MuhammadSuryono/module-golang-server/http"

func FailureResponse(typeStatus string, message string, data interface{}) models.CommonResponse {
	return models.CommonResponse{
		Code:      mapStatusCode(typeStatus),
		IsSuccess: false,
		Message:   message,
		Data:      data,
	}
}

func mapStatusCode(typeStatus string) int {
	switch typeStatus {
	case SUCCESS_STATUS:
		return SUCCESS_CODE
	case BAD_REQUEST_STATUS:
		return BAD_REQUEST_CODE
	case NOT_FOUND_STATUS:
		return NOT_FOUND_CODE
	case METHOD_NOT_ALLOWED_STATUS:
		return METHOD_NOT_ALLOWED_CODE
	case INTERNAL_SERVER_ERROR_STATUS:
		return INTERNAL_SERVER_ERROR_CODE
	case UNATHORIZED_STATUS:
		return UNATHORIZED_CODE
	default:
		return INTERNAL_SERVER_ERROR_CODE
	}
}

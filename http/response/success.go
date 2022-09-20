package response

import models "github.com/MuhammadSuryono/module-golang-server/http"

func SuccessResponse(created bool, message string, data interface{}) models.CommonResponse {
	var code = 200
	if created {
		code = 201
	}
	return models.CommonResponse{
		Code:      code,
		IsSuccess: true,
		Message:   message,
		Data:      data,
	}
}

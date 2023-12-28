package controllers

import "dating-app/models"

func BaseResponse(code int, message string, data interface{}) interface{} {
	return models.BaseResponseData{Code: code, Message: message, Data: data}
}

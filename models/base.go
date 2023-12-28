package models

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BaseResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

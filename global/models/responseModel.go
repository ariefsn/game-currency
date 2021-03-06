package models

type ResponseModel struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewResponseModel() *ResponseModel {
	return new(ResponseModel)
}

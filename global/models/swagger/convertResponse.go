package swagger

type ResponseConvert struct {
	Success bool    `json:"success" example:"true"`
	Data    float64 `json:"data" example:"0"`
	Message string  `json:"message" example:"some error message"`
}

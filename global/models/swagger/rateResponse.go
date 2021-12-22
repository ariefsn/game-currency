package swagger

type ResponseRate struct {
	Success bool   `json:"success" example:"true"`
	Data    Map    `json:"data" swaggertype:"object,string" example:"id:string,toIdCurrency:string,fromIdCurrency:string,rate:float64,createdAt:datetime,updatedAt:datetime"`
	Message string `json:"message" example:"some error message"`
}

type ResponseRateList struct {
	Success bool `json:"success" example:"true"`
	Data    ListMap
	Message string `json:"message" example:""`
}

package swagger

type ResponseCurrency struct {
	Success bool   `json:"success" example:"true"`
	Data    Map    `json:"data" swaggertype:"object,string" example:"id:string,name:string,createdAt:datetime,updatedAt:datetime"`
	Message string `json:"message" example:"some error message"`
}

type ResponseCurrencyList struct {
	Success bool `json:"success" example:"true"`
	Data    ListMap
	Message string `json:"message" example:""`
}

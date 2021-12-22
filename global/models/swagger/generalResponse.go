package swagger

type Map map[string]interface{}

type ListString struct {
	Items []interface{} `json:"items" swaggertype:"array,string"`
	Total int           `json:"total" example:"0"`
}

type ListMap struct {
	Items []interface{} `json:"items" swaggertype:"array,primitive,object"`
	Total int           `json:"total" example:"0"`
}

type ResponseError struct {
	Success bool        `json:"success" example:"false"`
	Data    interface{} `json:"data" swaggertype:"object"`
	Message string      `json:"message" example:"some error message"`
}

type ResponseSuccessText struct {
	Success bool   `json:"success" example:"true"`
	Data    string `json:"data" example:"some data"`
	Message string `json:"message" example:""`
}

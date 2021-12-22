package swagger

type PayloadConvert struct {
	FromIdCurrency string  `json:"fromIdCurrency" example:"string"`
	ToIdCurrency   string  `json:"toIdCurrency" example:"string"`
	Amount         float64 `json:"amount" example:"10"`
}

package swagger

type PayloadRate struct {
	FromIdCurrency string  `json:"fromIdCurrency" example:"string"`
	ToIdCurrency   string  `json:"toIdCurrency" example:"string"`
	Rate           float64 `json:"rate" example:"10"`
}

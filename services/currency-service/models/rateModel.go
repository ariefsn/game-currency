package models

import (
	"time"

	"ariefsn/game-currency/global/helper"

	"github.com/google/uuid"
)

type RateModel struct {
	Id             string    `json:"id" gorm:"type:varchar(36)"`
	FromIdCurrency string    `json:"fromIdCurrency" gorm:"type:varchar(36);column:fromIdCurrency"`
	ToIdCurrency   string    `json:"toIdCurrency" gorm:"type:varchar(36);column:toIdCurrency"`
	Rate           float64   `json:"rate"`
	CreatedAt      time.Time `json:"createdAt" gorm:"type:datetime;column:createdAt"`
	UpdatedAt      time.Time `json:"updatedAt" gorm:"type:datetime;column:updatedAt"`
}

func NewRateModel(from, to string, rate float64) *RateModel {
	m := new(RateModel)

	m.Id = uuid.NewString()
	m.FromIdCurrency = from
	m.ToIdCurrency = to
	m.Rate = rate
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	return m
}

func (m *RateModel) UpdateTime() *RateModel {
	m.UpdatedAt = time.Now()

	return m
}

func (m *RateModel) TableName() string {
	return "rates"
}

func (m *RateModel) HideFields(fields ...string) map[string]interface{} {
	return helper.HideFields(m, fields...)
}

func (m *RateModel) Convert(from, to string, amount float64) float64 {
	if m.FromIdCurrency == from && m.ToIdCurrency == to {
		return amount * m.Rate
	} else if m.ToIdCurrency == from && m.FromIdCurrency == to {
		return amount / m.Rate
	}

	return 0
}

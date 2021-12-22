package models

import (
	"time"

	"ariefsn/game-currency/global/helper"

	"github.com/google/uuid"
)

type CurrencyModel struct {
	Id        string    `json:"id" gorm:"type:varchar(36)"`
	Name      string    `json:"name" gorm:"type:varchar(100)"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:datetime;column:createdAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:datetime;column:updatedAt"`
}

func NewCurrencyModel(name string) *CurrencyModel {
	m := new(CurrencyModel)

	m.Id = uuid.NewString()
	m.Name = name
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	return m
}

func (m *CurrencyModel) UpdateTime() *CurrencyModel {
	m.UpdatedAt = time.Now()

	return m
}

func (m *CurrencyModel) TableName() string {
	return "currencies"
}

func (m *CurrencyModel) HideFields(fields ...string) map[string]interface{} {
	return helper.HideFields(m, fields...)
}

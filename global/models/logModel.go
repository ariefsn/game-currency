package models

import (
	"time"

	"github.com/google/uuid"
)

type LogModel struct {
	Id          string                 `json:"id"`
	ServiceName string                 `json:"serviceName" gorm:"column:serviceName"`
	Extras      map[string]interface{} `json:"extras" gorm:"column:extras"`
	Location    string                 `json:"location"`
	Description string                 `json:"description"`
	Status      int                    `json:"status"`
	CreatedAt   time.Time              `json:"createdAt" gorm:"column:createdAt"`
}

func NewLogModel() *LogModel {
	m := new(LogModel)

	m.Id = uuid.NewString()
	m.CreatedAt = time.Now()

	return m
}

func (m *LogModel) TableName() string {
	return "logs"
}

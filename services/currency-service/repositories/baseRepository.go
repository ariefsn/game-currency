package repositories

import (
	"ariefsn/game-currency/global/helper"

	"gorm.io/gorm"
)

type BaseRepository struct{}

var Msg = helper.NewMessage()

func DB() *gorm.DB {
	return helper.GetDB()
}

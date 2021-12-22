package helper

import (
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=%s", GetEnv().Db.User, GetEnv().Db.Password, GetEnv().Db.Host, GetEnv().Db.Name, GetEnv().Db.Timezone)

	var err error

	db, err = gorm.Open(mysql.Open(conn), &gorm.Config{})

	return err
}

func SetDB(newDb *gorm.DB) {
	db = newDb
}

func GetDB() *gorm.DB {
	return db
}

func MockDb() sqlmock.Sqlmock {
	sqlDb, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	db, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDb,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	SetDB(db)

	return mock
}

package repositories_test

import (
	"ariefsn/game-currency/global/helper"
	"ariefsn/game-currency/services/currency-service/models"
	"ariefsn/game-currency/services/currency-service/repositories"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRateRepository(t *testing.T) {

	var mock sqlmock.Sqlmock
	var currencyIds []string
	var insertedId string
	// var err error

	t.Run("Insert", func(t *testing.T) {
		mock = helper.MockDb()

		c := repositories.NewCurrencyRepository()
		r := repositories.NewRateRepository()

		dataCurrencies := []models.CurrencyModel{
			{
				Id:        uuid.NewString(),
				Name:      "Curr From",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Id:        uuid.NewString(),
				Name:      "Curr To",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		for _, curr := range dataCurrencies {
			sqlInsert := fmt.Sprintf("INSERT INTO `%s` (`id`,`name`,`createdAt`,`updatedAt`) VALUES (?,?,?,?)", curr.TableName())

			mock.ExpectExec(sqlInsert).WithArgs(curr.Id, curr.Name, curr.CreatedAt, curr.UpdatedAt).WillReturnResult(sqlmock.NewResult(1, 1))

			currId, err := c.Insert(curr)

			assert.Nil(t, err)
			assert.Equal(t, currId, curr.Id)

			currencyIds = append(currencyIds, insertedId)
		}

		dataRate := models.RateModel{
			Id:             uuid.NewString(),
			FromIdCurrency: dataCurrencies[0].Id,
			ToIdCurrency:   dataCurrencies[1].Id,
			Rate:           29,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		sqlFind := fmt.Sprintf("SELECT * FROM `%s` WHERE id = ? ORDER BY `currencies`.`id` LIMIT 1", dataCurrencies[0].TableName())
		sqlInsert := fmt.Sprintf("INSERT INTO `%s` (`id`,`fromIdCurrency`,`toIdCurrency`,`rate`,`createdAt`,`updatedAt`) VALUES (?,?,?,?,?,?)", dataRate.TableName())

		for _, curr := range dataCurrencies {
			mock.ExpectQuery(sqlFind).WithArgs(curr.Id).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "createdAt", "updatedAt"}).AddRow(curr.Id, curr.Name, curr.CreatedAt, curr.UpdatedAt))
		}
		mock.ExpectExec(sqlInsert).WithArgs(dataRate.Id, dataRate.FromIdCurrency, dataRate.ToIdCurrency, dataRate.Rate, dataRate.CreatedAt, dataRate.UpdatedAt).WillReturnResult(sqlmock.NewResult(1, 1))

		insertedId, err := r.Insert(dataRate)

		assert.Nil(t, err)
		assert.Equal(t, insertedId, dataRate.Id)

	})

	t.Run("Gets", func(t *testing.T) {
		r := repositories.NewRateRepository()

		data := models.RateModel{
			Id: insertedId,
		}

		sqlSelect := fmt.Sprintf("SELECT `id`,`fromIdCurrency`,`toIdCurrency`,`rate`,`createdAt`,`updatedAt` FROM `%s` LIMIT 100", data.TableName())
		sqlSelectCount := fmt.Sprintf("SELECT count(*) FROM `%s` LIMIT 100", data.TableName())

		mock.ExpectQuery(sqlSelect).WillReturnRows(sqlmock.NewRows([]string{"id", "fromIdCurrency", "toIdCurrency", "rate", "createdAt", "updatedAt"}).AddRow(data.Id, data.FromIdCurrency, data.ToIdCurrency, data.Rate, data.CreatedAt, data.UpdatedAt))
		mock.ExpectQuery(sqlSelectCount).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

		rows, total, err := r.Gets(0, 0, "")

		assert.Nil(t, err)
		assert.Equal(t, len(rows), 1)
		assert.Equal(t, total, 1)
	})

}

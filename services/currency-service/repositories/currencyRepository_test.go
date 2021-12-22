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

func TestCurrencyRepository(t *testing.T) {

	var mock sqlmock.Sqlmock
	var insertedId string
	var err error

	t.Run("Insert", func(t *testing.T) {
		mock = helper.MockDb()

		r := repositories.NewCurrencyRepository()

		data := models.CurrencyModel{
			Id:        uuid.NewString(),
			Name:      "Test Currency",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		sqlInsert := fmt.Sprintf("INSERT INTO `%s` (`id`,`name`,`createdAt`,`updatedAt`) VALUES (?,?,?,?)", data.TableName())

		mock.ExpectExec(sqlInsert).WithArgs(data.Id, data.Name, data.CreatedAt, data.UpdatedAt).WillReturnResult(sqlmock.NewResult(1, 1))

		insertedId, err = r.Insert(data)

		assert.Nil(t, err)
		assert.Equal(t, insertedId, data.Id)
	})

	t.Run("Gets", func(t *testing.T) {
		r := repositories.NewCurrencyRepository()

		data := models.CurrencyModel{
			Id: insertedId,
		}

		sqlSelect := fmt.Sprintf("SELECT `id`,`name`,`createdAt`,`updatedAt` FROM `%s` LIMIT 100", data.TableName())
		sqlSelectCount := fmt.Sprintf("SELECT count(*) FROM `%s` LIMIT 100", data.TableName())

		mock.ExpectQuery(sqlSelect).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "createdAt", "updatedAt"}).AddRow(data.Id, data.Name, data.CreatedAt, data.UpdatedAt))
		mock.ExpectQuery(sqlSelectCount).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

		rows, total, err := r.Gets(0, 0, "")

		assert.Nil(t, err)
		assert.Equal(t, len(rows), 1)
		assert.Equal(t, total, 1)
	})

	t.Run("Update", func(t *testing.T) {
		r := repositories.NewCurrencyRepository()

		data := models.CurrencyModel{
			Id:        insertedId,
			Name:      "Test Currency Update",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		sqlFind := fmt.Sprintf("SELECT * FROM `%s` WHERE id = ? ORDER BY `currencies`.`id` LIMIT 1", data.TableName())
		sqlUpdate := fmt.Sprintf("UPDATE `%s` SET `name`=?,`createdAt`=?,`updatedAt`=? WHERE id = ? AND `id` = ?", data.TableName())

		mock.ExpectQuery(sqlFind).WithArgs(data.Id).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "createdAt", "updatedAt"}).AddRow(data.Id, data.Name, data.CreatedAt, data.UpdatedAt))
		mock.ExpectExec(sqlUpdate).WithArgs(data.Name, sqlmock.AnyArg(), sqlmock.AnyArg(), data.Id, data.Id).WillReturnResult(sqlmock.NewResult(1, 1))

		updatedId, err := r.Update(insertedId, data)

		assert.Nil(t, err)
		assert.Equal(t, updatedId, data.Id)
	})

	t.Run("Delete", func(t *testing.T) {
		r := repositories.NewCurrencyRepository()

		data := models.CurrencyModel{
			Id: insertedId,
		}

		sqlFind := fmt.Sprintf("SELECT * FROM `%s` WHERE id = ? ORDER BY `currencies`.`id` LIMIT 1", data.TableName())
		sqlDelete := fmt.Sprintf("DELETE FROM `%s` WHERE id = ?", data.TableName())

		mock.ExpectQuery(sqlFind).WithArgs(data.Id).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "createdAt", "updatedAt"}).AddRow(data.Id, data.Name, data.CreatedAt, data.UpdatedAt))
		mock.ExpectExec(sqlDelete).WithArgs(data.Id).WillReturnResult(sqlmock.NewResult(1, 1))

		deletedId, err := r.Delete(insertedId)

		assert.Nil(t, err)
		assert.Equal(t, deletedId, data.Id)
	})

}

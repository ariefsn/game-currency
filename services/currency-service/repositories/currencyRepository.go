package repositories

import (
	"errors"

	"ariefsn/game-currency/global/helper"
	gmod "ariefsn/game-currency/global/models"
	"ariefsn/game-currency/services/currency-service/models"

	"github.com/google/uuid"
)

type CurrencyRepository interface {
	// GetById Get data by id. Return currency and an error.
	GetById(id string) (*models.CurrencyModel, error)
	// Gets Get list of data with skip, limit, and where for filtering. Return list, total list, and an error
	Gets(skip int, limit int, orders string, wheres ...gmod.FilterModel) ([]models.CurrencyModel, int, error)
	// Insert Create new data. Return inserted id, and an error.
	Insert(data models.CurrencyModel) (string, error)
	// Update Update existing data. Return updated id, and an error.
	Update(id string, data models.CurrencyModel) (string, error)
	// Delete Delete existing data. Return updated id, and an error.
	Delete(id string) (string, error)
}

type currencyRepository struct{}

func NewCurrencyRepository() CurrencyRepository {
	return new(currencyRepository)
}

// GetById Get data by id. Return currency and an error.
func (c *currencyRepository) GetById(id string) (*models.CurrencyModel, error) {
	res := models.CurrencyModel{}

	_db := DB().Table(res.TableName()).Where("id = ?", id).First(&res)

	if res.Id == "" {
		return nil, errors.New(Msg.DataNotFound)
	}

	return &res, _db.Error
}

// Gets Get list of data with skip, limit, and where for filtering. Return list, total list, and an error
func (c *currencyRepository) Gets(skip int, limit int, orders string, wheres ...gmod.FilterModel) ([]models.CurrencyModel, int, error) {
	mod := models.CurrencyModel{}
	res := []models.CurrencyModel{}

	qry := DB().Select("id", "name", "createdAt", "updatedAt").Table(mod.TableName())

	if wheres != nil {
		w, v := helper.BuildWhereClause(wheres...)
		qry = qry.Where(w, v...)
	}

	if skip != 0 {
		qry = qry.Offset(skip)
	}

	if limit == 0 {
		limit = 100
	}

	if orders != "" {
		qry = qry.Order(orders)
	}

	_db := qry.Limit(limit).Find(&res)

	if _db.Error != nil {
		return res, 0, _db.Error
	}

	total := helper.CountAllRows(qry)

	return res, total, nil
}

// Insert Create new data. Return inserted id, and an error.
func (c *currencyRepository) Insert(data models.CurrencyModel) (string, error) {
	if data.Id == "" {
		data.Id = uuid.NewString()
	}

	existsCurrency, err := c.GetById(data.Id)

	if err != nil {
		if err.Error() != Msg.DataNotFound {
			return "", err
		}
	}

	if existsCurrency != nil {
		if existsCurrency.Id == data.Id {
			return "", errors.New(Msg.DuplicateId)
		}
	}

	_db := DB().Table(data.TableName()).Create(&data)

	if _db.Error != nil {
		return "", _db.Error
	}

	return data.Id, nil
}

// Update Update existing data. Return updated id, and an error.
func (c *currencyRepository) Update(id string, data models.CurrencyModel) (string, error) {
	exists, err := c.GetById(id)

	if err != nil {
		return "", err
	}

	if exists.Id == id {
		data.Id = id
	}

	_db := DB().Table(data.TableName()).Where("id = ?", id).Updates(&data)

	if _db.Error != nil {
		return "", _db.Error
	}

	return id, nil
}

// Delete Delete existing data. Return updated id, and an error.
func (c *currencyRepository) Delete(id string) (string, error) {
	_, err := c.GetById(id)

	if err != nil {
		return "", err
	}

	feature := models.CurrencyModel{}

	_db := DB().Table(feature.TableName()).Where("id = ?", id).Delete(&feature)

	if _db.Error != nil {
		return "", _db.Error
	}

	return id, nil
}

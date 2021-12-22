package repositories

import (
	"errors"

	"ariefsn/game-currency/global/helper"
	gmod "ariefsn/game-currency/global/models"
	"ariefsn/game-currency/services/currency-service/models"

	"github.com/google/uuid"
)

type RateRepository interface {
	// GetByFromTo Get data by fromIdCurrency and toIdCurrency. Return currency and an error.
	GetByFromTo(from, to string) (*models.RateModel, error)
	// GetById Get data by id. Return currency and an error.
	GetById(id string) (*models.RateModel, error)
	// Gets Get list of data with skip, limit, and where for filtering. Return list, total list, and an error
	Gets(skip int, limit int, orders string, wheres ...gmod.FilterModel) ([]models.RateModel, int, error)
	// Insert Create new data. Return inserted id, and an error.
	Insert(data models.RateModel) (string, error)
	// Update Update existing data. Return updated id, and an error.
	Update(id string, data models.RateModel) (string, error)
	// Delete Delete existing data. Return updated id, and an error.
	Delete(id string) (string, error)
}

type rateRepository struct{}

func NewRateRepository() RateRepository {
	return new(rateRepository)
}

// GetByFromTo Get data by fromIdCurrency and toIdCurrency. Return currency and an error.
func (c *rateRepository) GetByFromTo(from, to string) (*models.RateModel, error) {
	res := &models.RateModel{}

	_db := DB().Table(res.TableName()).Where("(fromIdCurrency = ? AND toIdCurrency = ?) OR (fromIdCurrency = ? AND toIdCurrency = ?)", from, to, to, from).First(res)

	if res.Id == "" {
		return nil, errors.New(Msg.DataNotFound)
	}

	return res, _db.Error
}

// GetById Get data by id. Return currency and an error.
func (c *rateRepository) GetById(id string) (*models.RateModel, error) {
	res := models.RateModel{}

	_db := DB().Table(res.TableName()).Where("id = ?", id).First(&res)

	if res.Id == "" {
		return nil, errors.New(Msg.DataNotFound)
	}

	return &res, _db.Error
}

// Gets Get list of data with skip, limit, and where for filtering. Return list, total list, and an error
func (c *rateRepository) Gets(skip int, limit int, orders string, wheres ...gmod.FilterModel) ([]models.RateModel, int, error) {
	mod := models.RateModel{}
	res := []models.RateModel{}

	qry := DB().Select("id", "fromIdCurrency", "toIdCurrency", "rate", "createdAt", "updatedAt").Table(mod.TableName())

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
func (c *rateRepository) Insert(data models.RateModel) (string, error) {
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

	// check currency
	fromCurrency, _ := NewCurrencyRepository().GetById(data.FromIdCurrency)

	if fromCurrency == nil {
		return "", errors.New(Msg.FromCurrencyNotFound)
	}

	toCurrency, _ := NewCurrencyRepository().GetById(data.ToIdCurrency)

	if toCurrency == nil {
		return "", errors.New(Msg.ToCurrencyNotFound)
	}

	// check existing rate
	exists, _ := c.GetByFromTo(data.FromIdCurrency, data.ToIdCurrency)

	if exists != nil {
		return "", errors.New(Msg.Duplicate)
	}

	_db := DB().Table(data.TableName()).Create(&data)

	if _db.Error != nil {
		return "", _db.Error
	}

	return data.Id, nil
}

// Update Update existing data. Return updated id, and an error.
func (c *rateRepository) Update(id string, data models.RateModel) (string, error) {
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
func (c *rateRepository) Delete(id string) (string, error) {
	_, err := c.GetById(id)

	if err != nil {
		return "", err
	}

	feature := models.RateModel{}

	_db := DB().Table(feature.TableName()).Where("id = ?", id).Delete(&feature)

	if _db.Error != nil {
		return "", _db.Error
	}

	return id, nil
}

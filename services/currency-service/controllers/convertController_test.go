package controllers_test

import (
	"ariefsn/game-currency/global/helper"
	"ariefsn/game-currency/services/currency-service/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	t.Run("Convert", func(t *testing.T) {
		r := chi.NewRouter()
		controllers.NewConvertController().Register(r)
		controllers.NewCurrencyController().Register(r)
		controllers.NewRateController().Register(r)
		ts := httptest.NewServer(r)
		helper.InitEnv("../.env.example")
		helper.InitRender()
		helper.InitDB()
		defer ts.Close()

		// Insert Currencies
		currencyInsertedIds := []string{}
		currencyNames := []string{"From Currency", "ToCurrency"}
		for _, cn := range currencyNames {
			_, body := testRequest(t, ts, http.MethodPost, "/currency", buildPayload(map[string]interface{}{
				"name": cn,
			}))

			assert.True(t, body.Success)
			assert.NotEmpty(t, body.Data)
			currencyInsertedIds = append(currencyInsertedIds, body.Data.(string))
		}

		// Create Rate
		rateInsertedId := ""
		rateValue := 29
		_, body := testRequest(t, ts, http.MethodPost, "/rate", buildPayload(map[string]interface{}{
			"fromIdCurrency": currencyInsertedIds[1],
			"toIdCurrency":   currencyInsertedIds[0],
			"rate":           rateValue,
		}))

		assert.True(t, body.Success)
		assert.NotEmpty(t, body.Data)
		rateInsertedId = body.Data.(string)

		// Start Convert
		amount := 580
		convertPayloads := []map[string]interface{}{
			{
				"fromIdCurrency": currencyInsertedIds[0],
				"toIdCurrency":   currencyInsertedIds[1],
				"amount":         amount,
			},
			{
				"fromIdCurrency": currencyInsertedIds[1],
				"toIdCurrency":   currencyInsertedIds[0],
				"amount":         amount,
			},
		}

		convertAmountExpecteds := []float64{float64(amount) / float64(rateValue), float64(amount) * float64(rateValue)}

		for iCp, cp := range convertPayloads {
			_, body := testRequest(t, ts, http.MethodPost, "/convert", buildPayload(cp))

			assert.True(t, body.Success)
			assert.NotEmpty(t, body.Data)
			assert.Equal(t, convertAmountExpecteds[iCp], body.Data)
		}

		// Delete existing rate
		_, body = testRequest(t, ts, http.MethodDelete, "/rate/"+rateInsertedId, nil)
		assert.True(t, body.Success)
		assert.NotEmpty(t, body.Data)
		assert.Equal(t, rateInsertedId, body.Data)

		// Delete existing currencies
		for _, ci := range currencyInsertedIds {
			_, body = testRequest(t, ts, http.MethodDelete, "/currency/"+ci, nil)
			assert.True(t, body.Success)
			assert.NotEmpty(t, body.Data)
			assert.Equal(t, ci, body.Data)
		}
	})
}

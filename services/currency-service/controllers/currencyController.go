package controllers

import (
	"net/http"

	"ariefsn/game-currency/global/helper"
	gmod "ariefsn/game-currency/global/models"
	"ariefsn/game-currency/services/currency-service/models"
	rep "ariefsn/game-currency/services/currency-service/repositories"

	"github.com/go-chi/chi/v5"
)

var currencyRep = rep.NewCurrencyRepository()

type CurrencyController struct{}

func NewCurrencyController() *CurrencyController {
	return new(CurrencyController)
}

func (c *CurrencyController) Register(r *chi.Mux) *chi.Mux {
	r.Route("/currency", func(r chi.Router) {
		r.Get("/", c.Gets)
		r.Post("/", c.Create)
		r.Get("/{id}", c.Detail)
		r.Put("/{id}", c.Update)
		r.Delete("/{id}", c.Delete)
	})

	return r
}

// Gets godoc
// @Summary Currency list
// @Description Get currency list
// @Tags Currency
// @Accept json
// @Produce json
// @Param name query string false "Name"
// @Param skip query integer false "Skip"
// @Param limit query integer false "Limit"
// @Param orders query string false "Orders"
// @Success 200 {object} swagger.ResponseCurrencyList
// @Failure 400,404,500 {object} swagger.ResponseError
// @Router /currency [get]
// Handler for get currency list
func (c *CurrencyController) Gets(rw http.ResponseWriter, r *http.Request) {
	filters := []gmod.FilterModel{}
	filters = helper.AppendFilterField(r, "name", gmod.OprContains, filters)

	skip, limit := helper.GetSkipLimit(r)

	orders := helper.GetOrders(r)

	currencys, total, err := currencyRep.Gets(skip, limit, orders, filters...)

	if err != nil {
		RenderError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	if total == 0 {
		RenderSuccessList(rw, http.StatusOK, []interface{}{}, total)
		return
	}

	RenderSuccessList(rw, http.StatusOK, currencys, total)
}

// Create godoc
// @Summary Create currency
// @Description Create a new currency
// @Tags Currency
// @Accept json
// @Produce json
// @Param body body swagger.PayloadCurrency true "Payload"
// @Success 200 {object} swagger.ResponseSuccessText
// @Failure 400,404,500 {object} swagger.ResponseError
// @Router /currency [post]
// Handler for create new currency
func (c *CurrencyController) Create(rw http.ResponseWriter, r *http.Request) {
	payload := models.CurrencyModel{}

	if err := helper.ParsePayload(r, &payload); err != nil {
		RenderError(rw, http.StatusUnprocessableEntity, err.Error())
		return
	}

	idCurrency, err := currencyRep.Insert(payload)

	if err != nil {
		RenderError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	RenderSuccess(rw, http.StatusOK, idCurrency)
}

// Detail godoc
// @Summary Currency detail
// @Description Get currency detail
// @Tags Currency
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} swagger.ResponseCurrency
// @Failure 400,404,500 {object} swagger.ResponseError
// @Router /currency/{id} [get]
// Handler for get company currency detail
func (c *CurrencyController) Detail(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	currency, err := currencyRep.GetById(id)

	if err != nil {
		RenderError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	RenderSuccess(rw, http.StatusOK, currency)
}

// Update godoc
// @Summary Currency update
// @Description Update existing currency
// @Tags Currency
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param body body swagger.PayloadCurrency true "Payload"
// @Success 200 {object} swagger.ResponseSuccessText
// @Failure 400,404,500 {object} swagger.ResponseError
// @Router /currency/{id} [put]
// Handler for update currency
func (c *CurrencyController) Update(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	payload := models.CurrencyModel{}

	err := helper.ParsePayload(r, &payload)

	if err != nil {
		RenderError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	currency, err := currencyRep.Update(id, payload)

	if err != nil {
		RenderError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	RenderSuccess(rw, http.StatusOK, currency)
}

// Delete godoc
// @Summary Currency delete
// @Description Delete existing currency
// @Tags Currency
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} swagger.ResponseSuccessText
// @Failure 400,404,500 {object} swagger.ResponseError
// @Router /currency/{id} [delete]
// Handler for delete currency
func (c *CurrencyController) Delete(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	currency, err := currencyRep.Delete(id)

	if err != nil {
		RenderError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	RenderSuccess(rw, http.StatusOK, currency)
}

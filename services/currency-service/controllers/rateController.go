package controllers

import (
	"net/http"

	"ariefsn/game-currency/global/helper"
	gmod "ariefsn/game-currency/global/models"
	"ariefsn/game-currency/services/currency-service/models"
	rep "ariefsn/game-currency/services/currency-service/repositories"

	"github.com/go-chi/chi/v5"
)

var rateRep = rep.NewRateRepository()

type RateController struct{}

func NewRateController() *RateController {
	return new(RateController)
}

func (c *RateController) Register(r *chi.Mux) *chi.Mux {
	r.Route("/rate", func(r chi.Router) {
		r.Get("/", c.Gets)
		r.Post("/", c.Create)
		r.Get("/{id}", c.Detail)
		r.Put("/{id}", c.Update)
		r.Delete("/{id}", c.Delete)
	})

	return r
}

// Gets godoc
// @Summary Rate list
// @Description Get rate list
// @Tags Rate
// @Accept json
// @Produce json
// @Param skip query integer false "Skip"
// @Param limit query integer false "Limit"
// @Param orders query string false "Orders"
// @Success 200 {object} swagger.ResponseRateList
// @Failure 400,404,500 {object} swagger.ResponseError
// @Router /rate [get]
// Handler for get rate list
func (c *RateController) Gets(rw http.ResponseWriter, r *http.Request) {
	filters := []gmod.FilterModel{}

	skip, limit := helper.GetSkipLimit(r)

	orders := helper.GetOrders(r)

	rates, total, err := rateRep.Gets(skip, limit, orders, filters...)

	if err != nil {
		RenderError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	if total == 0 {
		RenderSuccessList(rw, http.StatusOK, []interface{}{}, total)
		return
	}

	RenderSuccessList(rw, http.StatusOK, rates, total)
}

// Create godoc
// @Summary Create rate
// @Description Create a new rate
// @Tags Rate
// @Accept json
// @Produce json
// @Param body body swagger.PayloadRate true "Payload"
// @Success 200 {object} swagger.ResponseSuccessText
// @Failure 400,404,500 {object} swagger.ResponseError
// @Router /rate [post]
// Handler for create new rate
func (c *RateController) Create(rw http.ResponseWriter, r *http.Request) {
	payload := models.RateModel{}

	if err := helper.ParsePayload(r, &payload); err != nil {
		RenderError(rw, http.StatusUnprocessableEntity, err.Error())
		return
	}

	idRate, err := rateRep.Insert(payload)

	if err != nil {
		RenderError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	RenderSuccess(rw, http.StatusOK, idRate)
}

// Detail godoc
// @Summary Rate detail
// @Description Get rate detail
// @Tags Rate
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} swagger.ResponseRate
// @Failure 400,404,500 {object} swagger.ResponseError
// @Router /rate/{id} [get]
// Handler for get company rate detail
func (c *RateController) Detail(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	rate, err := rateRep.GetById(id)

	if err != nil {
		RenderError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	RenderSuccess(rw, http.StatusOK, rate)
}

// Update godoc
// @Summary Rate update
// @Description Update existing rate
// @Tags Rate
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param body body swagger.PayloadRate true "Payload"
// @Success 200 {object} swagger.ResponseSuccessText
// @Failure 400,404,500 {object} swagger.ResponseError
// @Router /rate/{id} [put]
// Handler for update rate
func (c *RateController) Update(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	payload := models.RateModel{}

	err := helper.ParsePayload(r, &payload)

	if err != nil {
		RenderError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	rate, err := rateRep.Update(id, payload)

	if err != nil {
		RenderError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	RenderSuccess(rw, http.StatusOK, rate)
}

// Delete godoc
// @Summary Rate delete
// @Description Delete existing rate
// @Tags Rate
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} swagger.ResponseSuccessText
// @Failure 400,404,500 {object} swagger.ResponseError
// @Router /rate/{id} [delete]
// Handler for delete rate
func (c *RateController) Delete(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	rate, err := rateRep.Delete(id)

	if err != nil {
		RenderError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	RenderSuccess(rw, http.StatusOK, rate)
}

package controllers

import (
	"net/http"

	"ariefsn/game-currency/global/helper"
	swag "ariefsn/game-currency/global/models/swagger"

	"github.com/go-chi/chi/v5"
)

type ConvertController struct{}

func NewConvertController() *ConvertController {
	return new(ConvertController)
}

func (c *ConvertController) Register(r *chi.Mux) *chi.Mux {
	r.Route("/convert", func(r chi.Router) {
		r.Post("/", c.Start)
	})

	return r
}

// Start godoc
// @Summary Start convert
// @Description Start a new convertion
// @Tags Convert
// @Accept json
// @Produce json
// @Param body body swagger.PayloadConvert true "Payload"
// @Success 200 {object} swagger.ResponseConvert
// @Failure 400,404,500 {object} swagger.ResponseError
// @Router /convert [post]
// Handler for start convert
func (c *ConvertController) Start(rw http.ResponseWriter, r *http.Request) {
	payload := swag.PayloadConvert{}

	if err := helper.ParsePayload(r, &payload); err != nil {
		RenderError(rw, http.StatusUnprocessableEntity, err.Error())
		return
	}

	rate, err := rateRep.GetByFromTo(payload.FromIdCurrency, payload.ToIdCurrency)

	if err != nil {
		RenderError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	if rate == nil {
		RenderError(rw, http.StatusNotFound, Msg.DataNotFound)
		return
	}

	RenderSuccess(rw, http.StatusOK, rate.Convert(payload.FromIdCurrency, payload.ToIdCurrency, payload.Amount))
}

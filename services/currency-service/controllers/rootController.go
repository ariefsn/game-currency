package controllers

import (
	"fmt"
	"io"
	"net/http"

	"ariefsn/game-currency/global/helper"
	gmod "ariefsn/game-currency/global/models"

	"github.com/go-chi/chi/v5"
)

type RootController struct{}

func NewRootController() *RootController {
	return new(RootController)
}

var Msg = helper.NewMessage()

func RenderForward(w io.Writer, code int, data interface{}) {
	helper.Render().JSON(w, code, data)
}

func RenderSuccess(w io.Writer, code int, data interface{}) {
	m := gmod.NewResponseModel()

	m.Success = true
	m.Data = data

	helper.Render().JSON(w, code, m)
}

func RenderSuccessList(w io.Writer, code int, data interface{}, total int) {
	m := gmod.NewResponseModel()

	m.Success = true
	m.Data = map[string]interface{}{
		"items": data,
		"total": total,
	}

	helper.Render().JSON(w, code, m)
}

func RenderError(w io.Writer, code int, message string) {
	m := gmod.NewResponseModel()

	m.Success = false
	m.Message = message

	helper.Render().JSON(w, code, m)
}

func (c *RootController) Register(r *chi.Mux) *chi.Mux {
	r.Get("/", c.Ping)

	return r
}

// Ping godoc
// @Summary Ping provider service
// @Description Ping provider service
// @Tags Root
// @Accept json
// @Produce json
// @Success 200 {object} swagger.ResponseSuccessText
// @Failure 400,404,500 {object} swagger.ResponseError
// @Router / [get]
// Handler for ping provider service
func (c *RootController) Ping(rw http.ResponseWriter, r *http.Request) {
	RenderSuccess(rw, http.StatusOK, fmt.Sprintf("%v is running", helper.GetEnv().ServiceName))
}

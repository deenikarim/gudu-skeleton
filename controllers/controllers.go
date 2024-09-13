package controllers

import (
	"github.com/deenikarim/gudu"
	"myapp/data"
	"net/http"
)

type Handler struct {
	App    *gudu.Gudu
	Models data.Models
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {

	err := h.App.Render.RenderPage(w, r, "home.gohtml", nil, nil)

	if err != nil {
		h.App.ErrorLog.Println("error rendering home page:", err)
		return
	}
}

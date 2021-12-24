package handlers

import (
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/namnguyen191/goravel"
)

type Handlers struct {
	App *goravel.Goravel
}

func (h *Handlers) Home(rw http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(rw, r, "home", nil, nil)

	if err != nil {
		h.App.ErrorLog.Println("error rendering: ", err)
	}
}

func (h *Handlers) GoPage(rw http.ResponseWriter, r *http.Request) {
	err := h.App.Render.GoPage(rw, r, "home", nil)

	if err != nil {
		h.App.ErrorLog.Println("error rendering: ", err)
	}
}

func (h *Handlers) JetPage(rw http.ResponseWriter, r *http.Request) {
	err := h.App.Render.JetPage(rw, r, "jet-template", nil, nil)

	if err != nil {
		h.App.ErrorLog.Println("error rendering: ", err)
	}
}

func (h *Handlers) SessionTest(rw http.ResponseWriter, r *http.Request) {
	myData := "bar"

	h.App.Session.Put(r.Context(), "foo", myData)

	myValue := h.App.Session.GetString(r.Context(), "foo")

	vars := make(jet.VarMap)
	vars.Set("foo", myValue)

	err := h.App.Render.JetPage(rw, r, "sessions", vars, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering: ", err)
	}
}

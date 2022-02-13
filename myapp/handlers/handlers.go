package handlers

import (
	"fmt"
	"myapp/data"
	"net/http"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/namnguyen191/goravel"
)

type Handlers struct {
	App    *goravel.Goravel
	Models data.Models
}

func (h *Handlers) Home(rw http.ResponseWriter, r *http.Request) {
	defer h.App.LoadTime(time.Now())
	err := h.render(rw, r, "home", nil, nil)

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

func (h *Handlers) JSON(rw http.ResponseWriter, r *http.Request) {
	var payload struct {
		ID      int64    `json:"id,omitempty"`
		Name    string   `json:"name,omitempty"`
		Hobbies []string `json:"hobbies,omitempty"`
	}

	payload.ID = 10
	payload.Name = "Jack Jones"
	payload.Hobbies = []string{"karate", "tennis", "programming"}

	err := h.App.WriteJSON(rw, http.StatusOK, payload)
	if err != nil {
		h.App.ErrorLog.Println(err)
	}
}

func (h *Handlers) XML(rw http.ResponseWriter, r *http.Request) {
	type Payload struct {
		ID      int64    `xml:"id"`
		Name    string   `xml:"name"`
		Hobbies []string `xml:"hobbies>hobby"`
	}

	var payload Payload
	payload.ID = 10
	payload.Name = "John Smith"
	payload.Hobbies = []string{"karate", "tennis", "programming"}

	err := h.App.WriteXML(rw, http.StatusOK, payload)
	if err != nil {
		h.App.ErrorLog.Println(err)
	}
}

func (h *Handlers) DownloadFile(rw http.ResponseWriter, r *http.Request) {
	h.App.DownloadFile(rw, r, "./public/images/", "celeritas.jpg")
}

func (h *Handlers) TestCrypto(rw http.ResponseWriter, r *http.Request) {
	plainText := "Hello World!"
	fmt.Fprint(rw, "Unencrypted: "+plainText+"\n")
	encrypted, err := h.encrypt(plainText)
	if err != nil {
		h.App.ErrorLog.Println(err)
		h.App.Error500(rw, r)
		return
	}

	fmt.Fprint(rw, "Encrypted: "+encrypted+"\n")

	decrypted, err := h.decrypt(encrypted)
	if err != nil {
		h.App.ErrorLog.Println(err)
		h.App.Error500(rw, r)
		return
	}

	fmt.Fprint(rw, "Decrypted: "+decrypted+"\n")

}

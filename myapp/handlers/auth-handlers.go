package handlers

import "net/http"

func (h *Handlers) UserLogin(rw http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(rw, r, "login", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println(err)
	}
}

func (h *Handlers) PostUserLogin(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		rw.Write([]byte(err.Error()))
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := h.Models.Users.GetByEmail(email)
	if err != nil {
		rw.Write([]byte(err.Error()))
		return
	}

	matches, err := user.PasswordMatches(password)
	if err != nil {
		rw.Write([]byte("Error validating password"))
		return
	}

	if !matches {
		rw.Write([]byte("Invalid password"))
		return
	}

	h.App.Session.Put(r.Context(), "userID", user.ID)

	http.Redirect(rw, r, "/", http.StatusSeeOther)
}

func (h *Handlers) Logout(rw http.ResponseWriter, r *http.Request) {
	h.App.Session.RenewToken(r.Context())
	h.App.Session.Remove(r.Context(), "userID")
	http.Redirect(rw, r, "/users/login", http.StatusSeeOther)
}

package http_v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/giicoo/maratWebSite/models"
	"github.com/julienschmidt/httprouter"
	"github.com/noirbizarre/gonja"
	"github.com/sirupsen/logrus"
)

func (h *Handler) singUp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse ajax to json and send to Auth service
	logrus.Info(r.URL)

	body := r.Body
	defer body.Close()

	userToDB := models.User{}

	if err := json.NewDecoder(body).Decode(&userToDB); err != nil {
		logrus.Error(err)
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	user, err := h.services.AuthServices.SingUp(userToDB)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Service error", http.StatusInternalServerError)
		return
	}
	str := fmt.Sprint("Successful Add ", user.Login)
	w.Write([]byte(str))
}

func (h *Handler) singIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse user from ajax, send to auth service and check
	logrus.Info(r.URL)

	body := r.Body
	defer body.Close()

	user := models.User{}

	if err := json.NewDecoder(body).Decode(&user); err != nil {
		logrus.Error(err)
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	token, err := h.services.AuthServices.SingIn(user)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Invalid Form", http.StatusInternalServerError)
		return
	}
	// if we have token, we will save cookie
	ck := http.Cookie{
		Name:   "Auth",
		Value:  token,
		MaxAge: 3600,
	}

	http.SetCookie(w, &ck)
	w.Write([]byte("Successful login"))
}

func (h *Handler) singInUpPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Page singin/singup form
	logrus.Info(r.URL)

	var tpl = gonja.Must(gonja.FromFile("templates/logreg.html"))

	out, err := tpl.Execute(gonja.Context{"query": r.FormValue("query")})
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
	w.Write([]byte(out))
}

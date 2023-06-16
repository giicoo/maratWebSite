package http_v1

import (
	"encoding/json"
	"net/http"

	"github.com/giicoo/maratWebSite/models"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func (h *Handler) singUp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)

	body := r.Body
	defer body.Close()

	userToDB := models.User{}

	if err := json.NewDecoder(body).Decode(&userToDB); err != nil {
		logrus.Error(err)
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	if err := h.services.SingUp(userToDB); err != nil {
		logrus.Error(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Successful Add"))
}

func (h *Handler) singIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)

	body := r.Body
	defer body.Close()

	user := models.User{}

	if err := json.NewDecoder(body).Decode(&user); err != nil {
		logrus.Error(err)
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	token, err := h.services.SingIn(user)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Invalid Form", http.StatusInternalServerError)
		return
	}

	ck := http.Cookie{
		Name:  "Auth",
		Value: token,
	}

	http.SetCookie(w, &ck)
	w.Write([]byte("Successful login"))
}

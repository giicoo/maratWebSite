package http_v1

import (
	"net/http"
	"net/url"

	"github.com/giicoo/maratWebSite/internal/service/auth"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func (h *Handler) CookieAuthorization(handlerFunc httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		// JWT Token in cookie "Auth"
		token, err := r.Cookie("Auth")
		if err != nil {
			logrus.Error(err)
			http.Redirect(w, r, "/sing", http.StatusTemporaryRedirect)
			return
		}

		// Check token
		user, err := auth.ParseJWT(token.Value)
		if err != nil {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			logrus.Error(err)
			return
		}

		// if all OK, go to need handler func
		r.URL.User = url.User(user)
		handlerFunc(w, r, ps)
	}
}

func (h *Handler) CookieAuthorizationAdmin(handlerFunc httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		// JWT Token in cookie "Auth"
		token, err := r.Cookie("Auth")
		if err != nil {
			logrus.Error(err)
			http.Redirect(w, r, "/sing", http.StatusTemporaryRedirect)
			return
		}

		// Check token
		user, err := auth.ParseJWT(token.Value)
		if err != nil {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			logrus.Error(err)
			return
		}
		if user != "admin" {
			http.Error(w, "You are not admin", http.StatusUnauthorized)
			return
		}
		// if all OK, go to need handler func
		r.URL.User = url.User(user)
		handlerFunc(w, r, ps)
	}
}

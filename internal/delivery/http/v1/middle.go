package http_v1

import (
	"net/http"
	"net/url"

	"github.com/giicoo/maratWebSite/internal/service/auth"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func (h *Handler) BasicAuth(hand httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// JWT Token in cookie "Auth"
		// TODO: send user in func
		token, err := r.Cookie("Auth")
		if err != nil {
			logrus.Error(err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		user, err := auth.ParseJWT(token.Value)
		if err == nil {
			r.URL.User = url.User(user)
			hand(w, r, ps)
			return
		}
		http.Error(w, "Invalid Token", http.StatusUnauthorized)
		logrus.Error(err)
	}
}

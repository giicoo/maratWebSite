package http_v1

import (
	"net/http"

	"github.com/giicoo/maratWebSite/internal/service/auth"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func (h *Handler) BasicAuth(hand httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token, err := r.Cookie("Auth")
		if err != nil {
			logrus.Error(err)
			return
		}
		_, err = auth.ParseJWT(token.Value)
		if err == nil {
			hand(w, r, ps)
			return
		}
		logrus.Error(err)
	}
}

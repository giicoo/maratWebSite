package http_v1

import (
	"net/http"

	"github.com/giicoo/maratWebSite/internal/service"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitHandlers() http.Handler {
	r := httprouter.New()
	r.POST("/singup", h.singUp)
	r.POST("/singin", h.singIn)
	r.GET("/sing", h.sing)
	r.GET("/", h.BasicAuth(h.index))
	return r
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)
	w.Write([]byte("Home"))
}

func (h *Handler) sing(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)
	w.Write([]byte("Singin form"))
}

package http_v1

import (
	"net/http"

	"github.com/giicoo/maratWebSite/internal/service"
	"github.com/julienschmidt/httprouter"
	"github.com/noirbizarre/gonja"
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

	r.GET("/", h.BasicAuth(h.index))

	// auth
	r.POST("/singup", h.singUp)
	r.POST("/singin", h.singIn)
	r.GET("/sing", h.sing)

	// words
	r.POST("/add-word", h.addWord)
	r.POST("/get-words", h.getWords)

	r.ServeFiles("/templates/*filepath", http.Dir("templates"))
	return r
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)
	w.Write([]byte("Home"))
}

func (h *Handler) sing(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)
	var tpl = gonja.Must(gonja.FromFile("templates/logreg.html"))

	out, err := tpl.Execute(gonja.Context{"query": r.FormValue("query")})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte(out))

	w.Write([]byte("Singin form"))
}

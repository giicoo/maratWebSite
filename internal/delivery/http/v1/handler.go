package http_v1

import (
	"encoding/json"
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

	r.GET("/", h.BasicAuth(h.index))

	// auth
	r.POST("/singup", h.singUp)
	r.POST("/singin", h.singIn)
	r.GET("/sing", h.sing)

	// words
	r.POST("/add-word", h.addWord)
	r.POST("/get-words", h.getWords)
	r.POST("/debug", h.bug)

	//test
	r.GET("/test", h.testIndex)
	r.POST("/check-test", h.checkTest)

	r.ServeFiles("/templates/*filepath", http.Dir("templates"))
	return r
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)
	w.Write([]byte("Home"))
}

func (h *Handler) bug(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)
	words, _ := h.services.GetWordsForTest()
	jsonValue, err := json.Marshal(words)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonValue)
}

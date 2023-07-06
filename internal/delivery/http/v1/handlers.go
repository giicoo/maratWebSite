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

	r.GET("/", h.CookieAuthorization(h.index))

	// auth
	r.POST("/singup", h.singUp)
	r.POST("/singin", h.singIn)
	r.GET("/sing", h.singInUpPage)

	// words
	r.POST("/add-word", h.addWord)
	r.POST("/get-words", h.getWords)

	//test
	r.GET("/test/:name", h.testPageByName)
	r.POST("/check-test/:test_name", h.CookieAuthorization(h.checkTest))
	r.POST("/get-words-for-test/:name", h.getWordsForTest)
	r.POST("/add-test", h.addTest)

	// static file
	r.ServeFiles("/templates/*filepath", http.Dir("templates"))

	return r
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)
	w.Write([]byte("Home"))
}

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

	r.GET("/", h.CookieAuthorization(h.index))

	// auth
	r.POST("/singup", h.singUp)
	r.POST("/singin", h.singIn)
	r.GET("/sing", h.singInUpPage)
	r.POST("/logout", h.logout)

	// words
	r.GET("/create-word", h.CookieAuthorizationAdmin(h.createWordPage))
	r.POST("/add-word", h.CookieAuthorizationAdmin(h.addWord))
	r.POST("/get-words", h.CookieAuthorization(h.getWords))
	r.POST("/get-words-by-names", h.CookieAuthorization(h.getWordsByNames))

	//test
	r.GET("/tests", h.CookieAuthorization(h.testsPage))
	r.GET("/test/:name", h.CookieAuthorization(h.testPageByName))
	r.POST("/check-test/:test_name", h.CookieAuthorization(h.checkTest))
	r.POST("/test/res-page/:test_name", h.CookieAuthorization(h.resPage))
	r.POST("/get-words-for-test/:name", h.CookieAuthorization(h.getWordsForTest))
	r.GET("/create-test", h.CookieAuthorizationAdmin(h.createTestPage))
	r.POST("/add-test", h.CookieAuthorizationAdmin(h.addTest))

	// static file
	r.ServeFiles("/templates/*filepath", http.Dir("templates"))

	return r
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)

	// create template
	tpl := gonja.Must(gonja.FromFile("/templates/index.html"))

	out, err := tpl.Execute(gonja.Context{"user": r.URL.User})
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
	w.Write([]byte(out))
}

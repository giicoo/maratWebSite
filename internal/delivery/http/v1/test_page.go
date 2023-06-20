package http_v1

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/noirbizarre/gonja"
	"github.com/sirupsen/logrus"
)

func (h *Handler) testIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)
	tpl := gonja.Must(gonja.FromFile("/templates/main.html"))

	words, err := h.services.GetWord()
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Server error", http.StatusInternalServerError)
	}

	out, err := tpl.Execute(gonja.Context{"words": words, "first": words[0], "last": words[len(words)-1]})
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
	w.Write([]byte(out))

}

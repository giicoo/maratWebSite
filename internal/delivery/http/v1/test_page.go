package http_v1

import (
	"encoding/json"
	"net/http"

	"github.com/giicoo/maratWebSite/models"
	"github.com/julienschmidt/httprouter"
	"github.com/noirbizarre/gonja"
	"github.com/sirupsen/logrus"
)

func (h *Handler) testIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Page test
	logrus.Info(r.URL)

	tpl := gonja.Must(gonja.FromFile("/templates/main.html"))

	words, err := h.services.GetWordsForTest()
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
	w.Write([]byte("Test Page"))
}

func (h *Handler) checkTest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)

	body := r.Body
	defer body.Close()

	words := []*models.WordDB{}

	err := json.NewDecoder(body).Decode(&words)
	if err != nil {
		logrus.Error("JSON", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	answers, err := h.services.CheckTest(words)
	if err != nil {
		logrus.Error("SERVICE", err)
		http.Error(w, "Service Error", http.StatusInternalServerError)
		return
	}

	jsonValue, err := json.Marshal(answers)
	if err != nil {
		logrus.Error("JSON ANSWER", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonValue)
}

func (h *Handler) getWordsForTest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)

	body := r.Body
	defer body.Close()

	words, err := h.services.GetWordsForTest()
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Server error", http.StatusInternalServerError)
	}

	jsonValue, err := json.Marshal(words)
	if err != nil {
		logrus.Error("JSON ANSWER", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonValue)
}

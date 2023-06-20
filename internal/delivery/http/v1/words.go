package http_v1

import (
	"encoding/json"
	"net/http"

	"github.com/giicoo/maratWebSite/models"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func (h *Handler) addWord(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)

	body := r.Body
	defer body.Close()

	word := models.WordDB{}

	if err := json.NewDecoder(body).Decode(&word); err != nil {
		logrus.Error(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.services.AddWord(word); err != nil {
		logrus.Error(err)
		http.Error(w, "Service Error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Successful Add"))
}

func (h *Handler) getWords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)

	words, err := h.services.GetWord()
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Service Error", http.StatusInternalServerError)
	}

	jsonValue, err := json.Marshal(words)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonValue)
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

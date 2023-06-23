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
		logrus.Error(err, body)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.services.WordsServices.AddWord(word); err != nil {
		logrus.Error(err)
		http.Error(w, "Service Error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Successful Add"))
}

func (h *Handler) getWords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)

	words, err := h.services.WordsServices.GetWord()
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

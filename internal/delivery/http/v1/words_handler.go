package http_v1

import (
	"encoding/json"
	"net/http"

	"github.com/giicoo/maratWebSite/models"
	"github.com/julienschmidt/httprouter"
	"github.com/noirbizarre/gonja"
	"github.com/sirupsen/logrus"
)

func (h *Handler) addWord(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)

	body := r.Body
	defer body.Close()

	// parse request
	word := models.Word{}

	if err := json.NewDecoder(body).Decode(&word); err != nil {
		logrus.Error(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// add word
	if err := h.services.WordsServices.AddWord(word); err != nil {
		logrus.Error(err)
		http.Error(w, "Service Error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Successful Add"))
}

func (h *Handler) deleteWord(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)

	body := r.Body
	defer body.Close()

	word := []models.Word{}
	if err := json.NewDecoder(body).Decode(&word); err != nil {
		logrus.Error(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := h.services.WordsServices.DeleteWord(word)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Service error", http.StatusInternalServerError)
		return
	}
}
func (h *Handler) deleteWordPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)

	// get words
	words, err := h.services.WordsServices.GetWords()
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Service Error", http.StatusInternalServerError)
	}
	// create template
	tpl := gonja.Must(gonja.FromFile("templates/deleteword.html"))

	out, err := tpl.Execute(gonja.Context{"user": r.URL.User, "words": words})
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
	w.Write([]byte(out))
}
func (h *Handler) getWords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)

	// get words
	words, err := h.services.WordsServices.GetWords()
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Service Error", http.StatusInternalServerError)
	}

	// create response
	jsonValue, err := json.Marshal(words)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonValue)
}

func (h *Handler) getWordsByNames(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)

	body := r.Body
	defer body.Close()

	// parse request
	words_data := []*models.Word{}

	if err := json.NewDecoder(body).Decode(&words_data); err != nil {
		logrus.Error(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// get words
	words, err := h.services.WordsServices.GetWordsByNames(words_data)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Service Error", http.StatusInternalServerError)
	}

	// create response
	jsonValue, err := json.Marshal(words)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonValue)
}

func (h *Handler) createWordPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)

	// create template
	tpl := gonja.Must(gonja.FromFile("templates/createword.html"))

	out, err := tpl.Execute(gonja.Context{"user": r.URL.User})
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
	w.Write([]byte(out))
}

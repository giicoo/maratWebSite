package http_v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/giicoo/maratWebSite/models"
	"github.com/julienschmidt/httprouter"
	"github.com/noirbizarre/gonja"
	"github.com/sirupsen/logrus"
)

func (h *Handler) testsPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)
	tests, err := h.services.TestServices.GetTests()
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Service error", http.StatusInternalServerError)
		return
	}
	// create template
	tpl := gonja.Must(gonja.FromFile("/templates/listtest.html"))
	out, err := tpl.Execute(gonja.Context{"tests": tests, "user": r.URL.User})
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(out))

}
func (h *Handler) testPageByName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Page test with name from Params
	logrus.Info(r.URL)

	// get need words
	words, err := h.services.TestServices.GetWordsForTest(ps.ByName("name"))
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Server error", http.StatusInternalServerError)
	}

	// create template
	tpl := gonja.Must(gonja.FromFile("/templates/test.html"))

	out, err := tpl.Execute(gonja.Context{"words": words, "first": words[0], "last": words[len(words)-1], "test_name": ps.ByName("name"), "user": r.URL.User})
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
	w.Write([]byte(out))
}

func (h *Handler) checkTest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse json, check words
	logrus.Info(r.URL)

	body := r.Body
	defer body.Close()

	// parse request
	words := []*models.Word{}

	err := json.NewDecoder(body).Decode(&words)
	if err != nil {
		logrus.Error("JSON", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// check test
	answers, err := h.services.TestServices.CheckTest(words, ps.ByName("test_name"), r.URL.User.Username())
	if err != nil {
		logrus.Error("SERVICE", err)
		http.Error(w, "Service Error", http.StatusInternalServerError)
		return
	}

	// create response
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
	// get words
	logrus.Info(r.URL)

	body := r.Body
	defer body.Close()

	// get words
	words, err := h.services.TestServices.GetWordsForTest(ps.ByName("name"))
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Server error", http.StatusInternalServerError)
	}

	// create response
	jsonValue, err := json.Marshal(words)
	if err != nil {
		logrus.Error("JSON ANSWER", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonValue)
}

func (h *Handler) addTest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// add test
	logrus.Info(r.URL)

	body := r.Body
	defer body.Close()

	// parse request
	test := models.Test{UsersResults: []*models.UserResult{}}

	err := json.NewDecoder(body).Decode(&test)
	if err != nil {
		logrus.Error("JSON", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// add test
	err = h.services.TestServices.AddTest(test)
	if err != nil {
		logrus.Error("SERVICE", err)
		http.Error(w, "Service error", http.StatusInternalServerError)
	}
	str := fmt.Sprint("Successful Add ", test.Name)
	w.Write([]byte(str))
}

func (h *Handler) resPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)

	body := r.Body
	defer body.Close()

	answers := []*models.CheckTestWord{}

	err := json.NewDecoder(body).Decode(&answers)
	if err != nil {
		logrus.Error("JSON", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// create template
	tpl := gonja.Must(gonja.FromFile("/templates/res.html"))

	out, err := tpl.Execute(gonja.Context{"answers": answers, "test_name": ps.ByName("test_name"), "user": r.URL.User})
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
	w.Write([]byte(out))
}

func (h *Handler) createTestPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)

	words, err := h.services.WordsServices.GetWord()
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Service error", http.StatusInternalServerError)
		return
	}

	// create template
	tpl := gonja.Must(gonja.FromFile("/templates/createtest.html"))

	out, err := tpl.Execute(gonja.Context{"words": words, "user": r.URL.User})
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
	w.Write([]byte(out))
}

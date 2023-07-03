package http_v1

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func (h *Handler) bug(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)
	words, _ := h.services.TestServices.GetWordsForTest(ps.ByName("name"))
	jsonValue, err := json.Marshal(words)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonValue)
}

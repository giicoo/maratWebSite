package http_v1

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func (h *Handler) StatisticsGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)

	body := r.Body
	defer body.Close()

	tests, err := h.services.TestServices.GetTests()
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Service error", http.StatusInternalServerError)
		return
	}

	err = h.services.StatisticsServices.CreateStatistics(tests)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Service error", http.StatusInternalServerError)
		return
	}

	file, err := os.ReadFile(h.cfg.STAT_PATH)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Service error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Add("Content-Disposition", "attachment")
	w.Write(file)
}

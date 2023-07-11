package http_v1

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func (h *Handler) StatisticsGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)

	body := r.Body
	defer body.Close()

	// statistics := []*models.StatisticsExcel{{TestName: "test", Login: "tes", Percent: 10}}
	// detail_statistics := []*models.StatisticsUserExcel{{TestName: "test", Login: "tes", Percent: 10, CheckWordExcel: models.CheckWordExcel{WordExcel: models.WordExcel{Word: "t", Translate: "tt"}, Check: true, Right: "tt"}}}
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

	// err = h.services.StatisticsServices.WriteExcel(detail_statistics)
	// if err != nil {
	// 	logrus.Error(err)
	// 	http.Error(w, "Service error", http.StatusInternalServerError)
	// 	return
	// }
}

package service

import (
	"fmt"
	"reflect"

	"github.com/giicoo/maratWebSite/configs"
	"github.com/giicoo/maratWebSite/models"
	"github.com/plandem/xlsx"
	"github.com/plandem/xlsx/format/styles"
)

type StatisticsFunc interface {
	CreateStatistics(tests []*models.Test) error
}

type StatisticsService struct {
	cfg  *configs.Config
	path string
}

func (s *StatisticsService) CreateStatistics(tests []*models.Test) error {
	stats := []*models.StatisticsExcel{}
	detail_stats := []*models.CheckWordExcel{}
	for _, test := range tests {
		for _, res := range test.UsersResults {
			stat := &models.StatisticsExcel{TestName: test.Name}
			stat.Login = res.Login
			stat.Percent = res.Percent

			for _, wordres := range res.Res {
				detail_stat := &models.CheckWordExcel{}
				detail_stat.Login = res.Login
				detail_stat.TestName = stat.TestName
				detail_stat.Word = wordres.Word.Word
				detail_stat.Translate = wordres.Word.Translate
				detail_stat.Right = wordres.Right

				detail_stats = append(detail_stats, detail_stat)
			}

			stats = append(stats, stat)
		}
	}
	return WriteExcel(stats, detail_stats, s.cfg.STAT_PATH)
}

func WriteExcel(stats []*models.StatisticsExcel, detail_statistics []*models.CheckWordExcel, path string) error {
	xl := xlsx.New()
	defer xl.Close()

	//create a static sheet
	sheet := xl.AddSheet("Statistics")
	initMain(sheet)

	for j, row := range stats {
		j++
		e := reflect.ValueOf(row).Elem()

		for i := 0; i < e.NumField(); i++ {
			varName := e.Type().Field(i).Name
			varValue := e.Field(i).Interface()

			cell := sheet.Cell(i, j)

			if varName == "Login" {
				sheet := xl.SheetByName(fmt.Sprintf("%v", varValue))
				if sheet == nil {
					sheet = xl.AddSheet(fmt.Sprintf("%v", varValue))
				}
				sheet = initDetail(sheet)
				cell.SetValueWithHyperlink(varValue, fmt.Sprintf("#%v!A1", varValue))
			}
			cell.SetValue(varValue)

		}
	}

	// create detail static sheet
	rows := map[string]int{}
	for _, user := range detail_statistics {
		sheet := xl.SheetByName(user.Login)
		rows[user.Login]++

		e := reflect.ValueOf(user).Elem()

		for i := 0; i < e.NumField(); i++ {
			varName := e.Type().Field(i).Name
			varValue := e.Field(i).Interface()

			if varName != "Login" {
				cell := sheet.Cell(i, rows[user.Login])
				cell.SetValue(varValue)

				if varName == "Translate" {
					right := e.Field(i + 1).Interface()
					if varValue == right {
						cell.SetStyles(styles.New(
							styles.Fill.Type(styles.PatternTypeSolid),
							styles.Fill.Color("#81d41a"),
						))
					} else {
						cell.SetStyles(styles.New(
							styles.Fill.Type(styles.PatternTypeSolid),
							styles.Fill.Color("#ff0000"),
						))
					}
				}

			}
		}
	}

	return xl.SaveAs(path)
}

func initMain(sheet xlsx.Sheet) {
	stats := models.StatisticsExcel{TestName: "", Login: "", Percent: 0}
	e := reflect.ValueOf(&stats).Elem()
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varTag := e.Type().Field(i).Tag

		cell := sheet.Cell(i, 0)

		cell.SetValue(varName)
		cell.SetStyles(styles.New(
			styles.Fill.Type(styles.PatternTypeSolid),
			styles.Fill.Color(varTag.Get("color")),
		))
	}
}
func initDetail(sheet xlsx.Sheet) xlsx.Sheet {
	detail_statistics := models.CheckWordExcel{TestName: "", Word: "", Translate: "", Right: ""}
	e := reflect.ValueOf(&detail_statistics).Elem()
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varTag := e.Type().Field(i).Tag

		if varName != "Login" {
			cell := sheet.Cell(i, 0)

			cell.SetValue(varName)
			cell.SetStyles(styles.New(
				styles.Fill.Type(styles.PatternTypeSolid),
				styles.Fill.Color(varTag.Get("color")),
			))
		}
	}
	return sheet
}

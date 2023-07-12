package models

type StatisticsExcel struct {
	TestName string `color:"#ace466"`
	Login    string `color:"#ffff38"`
	Percent  int    `color:"#ffbf00"`
}

type CheckWordExcel struct {
	TestName  string `color:"#81d41a"`
	Word      string `color:"#bf819e"`
	Translate string `color:"#ffff6d"`
	Right     string `color:"#ffde59"`
	Login     string
}

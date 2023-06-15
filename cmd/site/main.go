package main

import (
	"github.com/giicoo/maratWebSite/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Server Start")
	err := app.Run()
	logrus.Fatal(err)
}

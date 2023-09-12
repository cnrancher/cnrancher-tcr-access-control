package main

import (
	"os"
	"runtime"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/cnrancher/tcr-access-control/commands"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&nested.Formatter{
		HideKeys:        false,
		TimestampFormat: "15:04:05", // hour, time, sec only
		FieldsOrder:     []string{},
	})
}

func main() {
	if runtime.GOOS == "windows" {
		logrus.Fatal("windows is not supported")
	}
	commands.Execute(os.Args[1:])
}

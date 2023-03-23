package pkg

import (
	"log"
	"os"

	"github.com/pterm/pterm"
)

func LogInfo(msg string, f func() error) {
	spinnerInfo, err := pterm.DefaultSpinner.Start(msg)
	if err != nil {
		log.Fatal(err)
	}

	if f == nil {
		spinnerInfo.Info()
		return
	}

	if err := f(); err != nil {
		LogFail(err.Error())
	}

	spinnerInfo.Info()
}

func LogFail(msg string) {
	spinnerExit, err := pterm.DefaultSpinner.Start(msg)
	if err != nil {
		log.Fatal(err)
	}

	spinnerExit.Fail()
	os.Exit(1)
}

func LogSuccess(msg string) {
	spinnerSuccess, err := pterm.DefaultSpinner.Start(msg)
	if err != nil {
		LogFail(err.Error())
	}

	spinnerSuccess.Success()
}

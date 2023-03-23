package main

import (
	"context"
	_ "embed"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/samuelsih/sucms/config"
	"github.com/samuelsih/sucms/pkg"
	"github.com/samuelsih/sucms/windows"
	"gopkg.in/yaml.v3"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go listenForCtrlC(ctx, stop)

	filename := "sucms-config.yml"

	args := os.Args
	if len(args) > 1 {
		filename = args[1]
	}

	pkg.LogInfo("Reading " + filename, nil)

	content, err := os.ReadFile(filename)
	if err != nil {
		pkg.LogFail(err.Error())
	}

	var config config.Raw
	if err := yaml.Unmarshal([]byte(content), &config); err != nil {
		log.Fatal(err)
	}

	if runtime.GOOS == "windows" {
		windows.RunScript(ctx, config)
	}

	pkg.LogSuccess("Success for generating " + config.ProjectName)
}

func listenForCtrlC(ctx context.Context, stop context.CancelFunc) {
	<-ctx.Done()
	stop()
	pkg.LogInfo("exiting...", nil)
	os.Exit(1)
}

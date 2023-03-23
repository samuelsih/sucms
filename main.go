package main

import (
	"context"
	_ "embed"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	info("Reading " + filename)

	content, err := os.ReadFile(filename)
	if err != nil {
		exit(err)
	}

	var config Config
	if err := yaml.Unmarshal([]byte(content), &config); err != nil {
		log.Fatal(err)
	}

	Generate(ctx, config)
}

func listenForCtrlC(ctx context.Context, stop context.CancelFunc) {
	<-ctx.Done()
	stop()
	exit("exiting...")
}

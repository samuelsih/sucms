package main

import (
	"context"
	"os/exec"
)

func Generate(ctx context.Context, config Config) {
	generateLaravel(ctx, config)
}

func generateLaravel(ctx context.Context, config Config) {
	info("Generating laravel file")

	name, procedure := GetCmd(ctx, CreateLaravelCmd, config)
	cmd := exec.CommandContext(ctx, name, procedure...)

	err := cmd.Start()
	if err != nil {
		exit(err)
	}

	err = cmd.Wait()
	if err != nil {
		exit(err)
	}
}

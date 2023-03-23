package main

import (
	"context"
	"runtime"
)

type Cmd int

const (
	CreateLaravelCmd Cmd = iota
	GenerateMigrationCmd
	EditMigrationCmd
)

func GetCmd(ctx context.Context, cmdType Cmd, config Config) (string, []string) {
	switch cmdType {
	case CreateLaravelCmd:
		if runtime.GOOS == "windows" {
			if !config.WantNewFolder {
				return "cmd", []string{"/C", "composer", "create-project", "laravel/laravel", "."}
			}

			return "cmd", []string{"/C", "composer", "create-project", "laravel/laravel", config.ProjectName}
		}

		if !config.WantNewFolder {
			return "bash", []string{"composer", "create-project", "laravel/laravel", "."}
		}

		return "bash", []string{"bash", "composer", "create-project", "laravel/laravel", config.ProjectName}

	default:
		exit("Bad command")
		return "", nil
	}
}

package windows

import (
	"github.com/samuelsih/sucms/config"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func generateLaravelProjectCMD(config config.Raw) (string, []string) {
	if config.WantNewFolder {
		return "cmd", []string{"/C", "composer", "create-project", "laravel/laravel", config.ProjectName}
	}

	return "cmd", []string{"/C", "composer", "create-project", "laravel/laravel", "."}
}

func generateMigrationCMD(config config.Raw) (string, [][]string) {
	var cmds [][]string

	for name := range config.Schemas {
		cmd := []string{"/C", "php", "artisan", "make:model", cases.Title(language.AmericanEnglish, cases.NoLower).String(name), "-m"}
		cmds = append(cmds, cmd)
	}

	return "cmd", cmds
}

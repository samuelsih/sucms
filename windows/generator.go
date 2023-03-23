package windows

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/samuelsih/sucms/config"
	"github.com/samuelsih/sucms/pkg"
)

func RunScript(ctx context.Context, raw config.Raw) {
	var extracted config.Extracted

	pkg.LogInfo("Generating Laravel File", func() error {
		return generateLaravel(ctx, raw)
	})

	pkg.LogInfo("Generating Model Based On Config", func() error {
		return generateModel(ctx, raw)
	})

	pkg.LogInfo("Generating Migration filename", func() error {
		extracted.RawMigrationFiles = getListMigrationName(raw)
		return nil
	})

	pkg.LogInfo("Find and Get Migration Filename", func() error {
		return nil
	})
}

func generateLaravel(ctx context.Context, config config.Raw) error {
	name, procedure := generateLaravelProjectCMD(config)
	cmd := exec.CommandContext(ctx, name, procedure...)

	err := cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

func generateModel(ctx context.Context, config config.Raw) error {
	name, procs := generateMigrationCMD(config)

	for _, proc := range procs {
		proc := proc //avoid race

		cmd := exec.CommandContext(ctx, name, proc...)
		if config.WantNewFolder {
			cmd.Dir = config.ProjectName
		}

		err := cmd.Start()
		if err != nil {
			return err
		}

		err = cmd.Wait()
		if err != nil {
			return err
		}
	}

	return nil
}

func getListMigrationName(config config.Raw) []string {
	var result []string

	for name := range config.Schemas {
		result = append(result, getTableName(name))
	}

	return result
}

func getTableName(s string) string {
	s = strings.ToLower(s)

	if strings.HasSuffix(s, "y") {
		s := s[:len(s)-1]
		s += "ies"
		return fmt.Sprintf("create_%s_table.php", s)
	}

	if !strings.HasSuffix(s, "s") {
		return fmt.Sprintf("create_%ss_table.php", s)
	}

	return fmt.Sprintf("create_%s_table.php", s)
}

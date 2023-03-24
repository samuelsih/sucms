package windows

import (
	"bufio"
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

	pkg.LogInfo("Generating Migration Based On Config", func() error {
		return generateMigration(ctx, raw)
	})

	pkg.LogInfo("Extracting Migration filename", func() error {
		extracted.ExpectedMigrationFilename = getListMigrationName(raw)
		return nil
	})

	pkg.LogInfo("Find and Get Migration Filename", func() error {
		fileLists, err := findMigrationfile(ctx, raw, extracted.ExpectedMigrationFilename)
		if err != nil {
			return err
		}

		extracted.MigrationFiles = fileLists
		return nil
	})

	pkg.LogInfo("Extracting Schemas", func() error {
		extracted.SchemasDefined = extractSchemas(raw)
		return nil
	})

	fmt.Println(extracted)
}

// Generating Laravel File
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

// Generating Migration Based On Config
func generateMigration(ctx context.Context, config config.Raw) error {
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

// Extracting Migration filename
func getListMigrationName(config config.Raw) []string {
	var result []string

	for name := range config.Schemas {
		result = append(result, getTableName(name))
	}

	return result
}

// Find and Get Migration Filename
func findMigrationfile(ctx context.Context, cfg config.Raw, filenames []string) ([]string, error) {
	name, proc := findMigrationFileCMD()

	cmd := exec.CommandContext(ctx, name, proc...)
	cmd.Dir = fmt.Sprintf("%s/database/migrations", cfg.ProjectName)

	if !cfg.WantNewFolder {
		cmd.Dir = "database/migrations"
	}

	r, _ := cmd.StdoutPipe()
	buf := bufio.NewScanner(r)
	var results []string

	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	for buf.Scan() {
		line := buf.Text()

		for _, filename := range filenames {
			if strings.Contains(line, filename) {
				splitted := strings.Split(line, " ")
				results = append(results, splitted[len(splitted)-1])
			}
		}
	}

	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	return results, nil
}

func extractSchemas(cfg config.Raw) config.SchemasDefined {
	configMapper := make(config.SchemasDefined)

	for name, value := range cfg.Schemas {
		newName := fmt.Sprintf("%s|%s", name, getTableName(name))
		configMapper[newName] = value
	}

	return configMapper
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

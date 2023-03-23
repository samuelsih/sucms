package config

type Schema map[string]string

type Raw struct {
	ProjectName   string            `yaml:"project_name"`
	WantNewFolder bool              `yaml:"new_folder"`
	Schemas       map[string]Schema `yaml:"schema"`
}

type Extracted struct {
	RawMigrationFiles []string
	MigrationFiles    []string
	Dir               string
}

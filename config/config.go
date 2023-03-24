package config

type Schema = map[string]string
type SchemasDefined = map[string]Schema

type Raw struct {
	ProjectName   string         `yaml:"project_name"`
	WantNewFolder bool           `yaml:"new_folder"`
	Schemas       SchemasDefined `yaml:"schema"`
}

type Extracted struct {
	ExpectedMigrationFilename []string
	MigrationFiles            []string
	SchemasDefined            SchemasDefined
}

package main

type Schema map[string]string

type Config struct {
	ProjectName   string            `yaml:"project_name"`
	WantNewFolder bool              `yaml:"new_folder"`
	Schemas       map[string]Schema `yaml:"schema"`
}

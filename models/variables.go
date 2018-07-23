package models

// Variable variables for the workflow
type Variable struct {
	Name   string `yaml:"name"`
	Type   string `yaml:"type"`
	Value  string `yaml:"value"`
	Format string `yaml:"format"`
}

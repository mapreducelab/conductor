package models

// Action represents action object
type Action struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`
	DisplayName string `yaml:"displayName"`
	Description string `yaml:"description"`
	Next        Task   `yaml:"next"`
	ShellScript string `yaml:"shellScript"`
	Output      string `yaml:"output"`
}

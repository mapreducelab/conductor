package models

// Workflow represents series for tasks
type Workflow struct {
	Variables   []Variable `yaml:"variables"`
	Rootdoc     string     `yaml:"rootdoc"`
	Type        string     `yaml:"type"`
	Name        string     `yaml:"name"`
	DisplayName string     `yaml:"displayName"`
	Description string     `yaml:"description"`
	Input       Input      `yaml:"input"`
	Start       Task       `yaml:"start"`
	Actions     []Action   `yaml:"actions"`
}

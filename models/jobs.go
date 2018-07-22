package models

// Job model
type Job struct {
	OsEnv    OsEnv         `yaml:"osEnv"`
	Project  Project       `yaml:"project"`
	Logging  []LoggingItem `yaml:",flow"`
	Workflow Workflow      `yaml:"workflow"`
}

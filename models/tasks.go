package models

// Task represents start object
type Task struct {
	Parallelism int      `yaml:"parallelism"`
	Multithread bool     `yaml:"multithread"`
	Subscribers []string `yaml:"subscribers"`
}

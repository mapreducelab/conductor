package models

// LoggingItem model
type LoggingItem struct {
	Name string `yaml:"name"`
	File File   `yaml:"file"`
	Mail Mail   `yaml:"mail"`
}

// File model
type File struct {
	Name      string `yaml:"name"`
	Directory string `yaml:"directory"`
	Extension string `yaml:"extension"`
	Pattern   string `yaml:"pattern"`
}

// Mail model
type Mail struct {
	Subject string `yaml:"subject"`
	List    string `yaml:"list"`
	Cmd     string `yaml:"cmd"`
}

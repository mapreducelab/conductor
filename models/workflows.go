package models

// Workflow model
type Workflow struct {
	Host      string   `yaml:"host"`
	Staging   string   `yaml:"staging"`
	Test      Test     `yaml:"test"`
	Urls      []string `yaml:",flow"`
	Directory string   `yaml:"directory"`
	Cursor    Cursor   `yaml:"cursor"`
	Timestamp string   `yaml:"timestamp"`
	Tasks     []Task   `yaml:",flow"`
}

// Test model
type Test struct {
	Flag      bool     `yaml:"flag"`
	Functions []string `yaml:",flow"`
}

// Cursor model
type Cursor struct {
	Range      []string `yaml:",flow"`
	Dateformat string   `yaml:"dateformat"`
	Format     string   `yaml:"format"`
	Increment  string   `yaml:"increment"`
	Inherit    bool     `yaml:"inherit"`
}

// Task model
type Task struct {
	Name  string      `yaml:"name"`
	Desc  string      `yaml:"desc"`
	Start []StartItem `yaml:",flow"`
	Steps []Step      `yaml:"steps"`
	Next  Next        `yaml:"next"`
}

// StartItem model
type StartItem struct {
	Order string   `yaml:"order"`
	Tasks []string `yaml:",flow"`
}

// Step model
type Step struct {
	Name string   `yaml:"name"`
	Desc string   `yaml:"desc"`
	Next StepNext `yaml:"next"`
	Call Call     `yaml:"call"`
}

// Next model
type Next struct {
	Success string `yaml:"success"`
	Failure string `yaml:"failure"`
}

// StepNext model
type StepNext struct {
	Success []SuccessItem `yaml:",flow"`
}

// SuccessItem model
type SuccessItem struct {
	Order int      `yaml:"order"`
	Tasks []string `yaml:",flow"`
}

// Call model
type Call struct {
	Action string `yaml:"action"`
	Thread int    `yaml:"thread"`
	Cursor Cursor `yaml:"cursor"`
	Exec   Exec   `yaml:"exec"`
}

// Exec model
type Exec struct {
	URL string   `yaml:"url"`
	Cmd []string `yaml:",flow"`
}

package models

// Blueprint represents blueprint for workflow or platform.
type Blueprint struct {
	Topology Topology `yaml:"topology"`
	Views    Views    `yaml:"views"`
}

// Topology model.
type Topology struct {
	Nodes []Node `yaml:"nodes"`
}

// Node model.
type Node struct {
	ID          string   `yaml:"id"`
	Network     Network  `yaml:"network"`
	Region      string   `yaml:"region"`
	Pod         string   `yaml:"pod"`
	Components  []string `yaml:"components"`
	Connections []string `yaml:"connections"`
}

// Network model.
type Network struct {
	Count        int    `yaml:"count"`
	Rootname     string `yaml:"rootname"`
	Domain       string `yaml:"domain"`
	ExposedPorts []Port `yaml:"exposed_ports"`
}

// Port model.
type Port struct {
	Port int    `yaml:"port"`
	Type string `yaml:"type"`
}

// Views model.
type Views struct {
	Platform Platform `yaml:"platform"`
}

// Platform model.
type Platform struct {
	Components []Component `yaml:"components"`
}

// Component model.
type Component struct {
	ID             string            `yaml:"id"`
	Automation     Automation        `"yaml:"automation"`
	Opscontext     Opscontext        `yaml:"opscontext"`
	Servicecontext map[string]string `yaml:"servicecontext"`
}

// Automation model.
type Automation struct {
	Type     string   `yaml:'yaml'`
	Artifact Artifact `yaml:"artifact"`
}

// Artifact model.
type Artifact struct {
	Name    string `yaml:"name"`
	Version string `yaml:'version'`
	Build   int    `yaml:"build"`
}

// Opscontext model.
type Opscontext struct {
	Storage Storage  `yaml:"storage"`
	Logging Logging  `yaml:"logging"`
	Secrets []Secret `yaml:"secrets"`
}

// Storage model.
type Storage struct {
	Drive Drive `yaml:"drive"`
}

// Drive model.
type Drive struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Location string `yaml:"location"`
}

// Logging model.
type Logging struct {
	Loglevel string `yaml:"loglevel"`
}

// Secrets model.
type Secret struct {
	Type   string `yaml:"type"`
	Secret string `yaml:"secret"`
}

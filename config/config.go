package config

var (
	// Etcd variable
	Etcd string

	config *configStruct
)

type configStruct struct {
	Etcd string `yaml:"etcd"`
}

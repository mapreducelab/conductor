package maestro

import (
	"errors"
)

// Maestro submits blueprint to key/value store.
type Maestro interface {
	Deploy(string, string) (bool, error)
	Undeploy(string, string) (bool, error)
}

type maestro struct{}

func (maestro) Deploy(pod string, blueprint string) (bool, error) {
	if pod == "" {
		return false, ErrEmptyPod
	}
	return true, nil
}

func (maestro) Undeploy(pod string, blueprint string) (bool, error) {
	if pod == "" {
		return false, ErrEmptyPod
	}
	return true, nil
}

// NewService placeholder for dependencies.
func NewService() Maestro {
	return &maestro{}
}

// ErrEmptyPod is returned when an pod string is empty.
var ErrEmptyPod = errors.New("Pod is empty string")

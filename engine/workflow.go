package engine

import "errors"

// Workflow engine processes job
type Workflow interface {
	ProcessJob(job string) error
}

type workflow struct{}

func (w *workflow) Workflow(job string) error {
	if job == "" {
		return errors.New("Job is empty")
	}
	return nil
}

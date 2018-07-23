package engine

import (
	"conductor/drivers"
	"conductor/models"
	"errors"
	"fmt"
	"strings"
)

// Workflow engine
type Workflow struct{}

// ProcessWorkflow creates a DIAG graph and process its components
func (w Workflow) ProcessWorkflow(workflow models.Workflow) error {
	if len(workflow.Actions) == 0 {
		return errors.New("Workflow is empty")
	}

	for _, action := range workflow.Actions {
		processAction(action)
	}

	return nil
}

func processAction(action models.Action) error {
	if action.Name == "" {
		return errors.New("action is empty")
	}
	// TODO dynamicly load drivers
	switch t := action.Type; strings.ToUpper(t) {
	case "SHELL":
		s := drivers.Shell{}
		res, _ := s.Deploy(action)
		fmt.Println(res)
	default:
		fmt.Printf("%v - wrong driver type", t)
	}

	fmt.Println(action)
	return nil
}

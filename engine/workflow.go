package engine

import (
	"conductor/models"
	"fmt"
	"strings"
)

// Workflow engine
type Workflow struct{}

// ProcessWorkflow creates a DIAG graph and process its components
func (w Workflow) ProcessWorkflow(workflow models.Blueprint) error {

	for _, component := range workflow.Views.Platform.Components {
		processAction(component)
	}

	return nil
}

func processAction(component models.Component) error {

	switch t := component.Automation.Type; strings.ToUpper(t) {
	case "SHELL":
		fmt.Println("result")
	default:
		fmt.Printf("%v - wrong driver type", t)
	}

	fmt.Println(component)
	return nil
}

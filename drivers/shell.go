package drivers

import (
	"conductor/models"
	"conductor/providers"
	"errors"
)

// Shell driver
type Shell struct {
	Host string
}

// Deploy func
func (s Shell) Deploy(action models.Action) (models.ActionResult, error) {
	if action.ShellScript == "" {
		return models.ActionResult{}, errors.New("action object does not have any commands to run")
	}

	result := models.ActionResult{}
	shellProvider := providers.Shell{}

	res, err := shellProvider.Exec(action.ShellScript)
	if err != nil {
		return models.ActionResult{}, err
	}
	result.Output = append(result.Output, res)

	return result, nil
}

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
func (s Shell) Deploy(call models.Call) (models.CallResult, error) {
	if len(call.Exec.Cmd) == 0 {
		return models.CallResult{}, errors.New("call object does not have any commands to run")
	}

	callCmds := call.Exec.Cmd
	result := models.CallResult{}
	shellProvider := providers.Shell{}

	for _, cmd := range callCmds {
		res, err := shellProvider.Exec(cmd)
		if err != nil {
			return models.CallResult{}, err
		}
		result.Output = append(result.Output, res)
	}

	return result, nil
}

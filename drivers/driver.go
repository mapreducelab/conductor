package drivers

import "conductor/models"

// Driver interface
type Driver interface {
	Deploy(models.Call) (models.CallResult, error)
}

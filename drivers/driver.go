package drivers

import "conductor/models"

// Driver interface
type Driver interface {
	Deploy(models.Action) (models.ActionResult, error)
}

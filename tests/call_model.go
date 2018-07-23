package tests

import "conductor/models"

// ActionTest model for testing
var ActionTest = models.Action{
	Name:        "shell701",
	Type:        "SHELL",
	DisplayName: "Print average rate shell701",
	Description: "Prints average BTC/USD rate in the canonical form",
	Next: models.Task{
		Parallelism: 1,
		Multithread: false,
	},
	ShellScript: "ls -ltr /tmp",
}

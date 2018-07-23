package tests

import "conductor/models"

// CallTest model for testing
var CallTest = models.Call{
	Action: "SHELL",
	Thread: 1,
	Cursor: models.Cursor{
		Range:      []string{"0"},
		Dateformat: "integer",
		Increment:  "1",
		Inherit:    false,
	},
	Exec: models.Exec{
		URL: "localhost",
		Cmd: []string{"ls -ltr /tmp", "echo ~/Code/conductor/conductor"},
	},
}

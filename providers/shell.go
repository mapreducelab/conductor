package providers

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

// Shell struct
type Shell struct{}

// Exec executes command with args
func (sp Shell) Exec(command string) (string, error) {
	if command == "" {
		return "", errors.New("provided command is empty string")
	}

	cmdName, args, err := parseCmd(command)
	if err != nil {
		return "", err
	}

	cmd := exec.Command(cmdName, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func parseCmd(command string) (string, []string, error) {
	if command == "" {
		return "", nil, errors.New("cdm is empty string")
	}

	cmds := strings.Split(command, " ")
	cmd := cmds[0]
	args := cmds[1:]

	return cmd, args, nil
}

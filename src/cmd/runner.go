package cmd

import (
	"fmt"
	"os/exec"
)

func (command InternalCommand)Execute() ([]byte, error) {

	terminal := exec.Command(command.Application, command.Args...)
	output, err := terminal.Output()

	return output, err
}
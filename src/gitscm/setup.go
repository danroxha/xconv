package gitscm

import (
	"runtime"
	"strings"

	"github.com/dannrocha/xconv/src/cmd"
)

func IsGitInstalled() bool {
	var command string

	if runtime.GOOS == "windows" {
		command = `where`		
	} else {
		command = `which`
	}

	binaryGit := cmd.InternalCommand{
		Application: command,
		Args: []string{
			`git`,
		},
	}

	_, err := binaryGit.Execute()

	return err == nil
}

func IsStageAreaEmpty() bool {
	command := cmd.InternalCommand{
		Application: `git`,
		Args: []string{
			`diff`,
			`--staged`,
		},
	}

	output, err := command.Execute()

	if err != nil {
		return err == nil
	}

	return strings.Replace(string(output), " ", "", -1) == ""
}

func IsGitRepository() bool {
	command := cmd.InternalCommand{
		Application: `git`,
		Args: []string{
			`status`,
		},
	}

	_, err := command.Execute()

	return err == nil
}
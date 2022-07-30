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
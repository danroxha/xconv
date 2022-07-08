package gitscm

import (
	"github.com/dannrocha/czen/src/cli"
)

func BackToCommit(commit string) error {
	reset := cmd.InternalCommand{
		Application: "git",
		Args: []string{
			"reset",
			"--hard",
			commit,
		},
	}

	_, err := reset.Execute()

	return err
}

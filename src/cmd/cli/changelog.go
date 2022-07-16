package cli

import (
	"github.com/dannrocha/czen/src/setup"
	"github.com/urfave/cli/v2"
)

func Changelog(c *cli.Context) error {
	scrip := setup.NewScript()

	for _, task := range scrip.Task {
		if task.Bind == CHANGELOG && task.Enable {
			if task.When == setup.BEFORE {
				task.Run()
			} else {
				defer task.Run()
			}
		}
	}
	return nil
}

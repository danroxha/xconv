package cli

import (
	"fmt"

	"github.com/dannrocha/czen/src/setup"
	"github.com/urfave/cli/v2"
)

func Init(c *cli.Context) error {

	script := setup.NewScript()

	for _, task := range script.Task {
		if task.Bind == INIT && task.Enable {
			if task.When == setup.BEFORE {
				task.Run()
			} else {
				defer task.Run()
			}
		}
	}

	fmt.Println("init not implemented")
	return nil
}

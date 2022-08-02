package cli

import (
	"fmt"
	"os"

	"github.com/dannrocha/xconv/src/setup"
	"github.com/urfave/cli/v2"
)

func Example(c *cli.Context) error {
	script := setup.NewScript()

	for _, task := range script.Task {
		if task.Bind == EXAMPLE && task.Enable {
			if task.When == setup.BEFORE {
				task.Run()
			} else {
				defer task.Run()
			}
		}
	}

	rule := setup.NewRule()

	profile, err := rule.FindCurrentProfileEnable()

	if err != nil {
		exception := setup.ExitCodeStardard["ActiveProfileNotFound"]
		fmt.Println(exception.Description)
		os.Exit(exception.ExitCode)
	}

	fmt.Println(profile.Example)

	return nil
}

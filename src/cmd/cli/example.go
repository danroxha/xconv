package cli

import (
	"fmt"

	"github.com/dannrocha/czen/src/setup"
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

	profile, errProf := rule.FindCurrentProfileEnable()

	if errProf != nil {
		panic(errProf)
	}

	fmt.Println(profile.Example)

	return nil
}

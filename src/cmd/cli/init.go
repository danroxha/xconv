package cli

import (
	"fmt"

	"github.com/dannrocha/czen/src/setup"
	"github.com/urfave/cli/v2"
)

func Init(c *cli.Context) error {

	script := setup.NewScript()

	for _, auto := range script.Automation {
		if auto.Bind == INIT && auto.Enable {
			if auto.When == setup.BEFORE {
				auto.Run()
			} else {
				defer auto.Run()
			}
		}
	}

	fmt.Println("init not implemented")
	return nil
}

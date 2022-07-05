package cli

import (
	"github.com/dannrocha/czen/src/setup"
	"github.com/urfave/cli/v2"
)


func Bump(c *cli.Context) error {
	
	scrip := setup.Script{}
	scrip.LoadScript()

	for _, auto := range scrip.Automation {
		if auto.Bind == BUMP && auto.Enable {
			if auto.When == setup.BEFORE {
				auto.Run()
			} else {
				defer auto.Run()
			}
		}
	}
	
	return nil
}

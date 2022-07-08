package bump

import (
	"github.com/dannrocha/czen/src/cli"
	"github.com/dannrocha/czen/src/setup"
)

func Execute(args ...string) error {
	scrip := setup.Script{}
	scrip.LoadScript()

	for _, auto := range scrip.Automation {
		if auto.Bind == cmd.BUMP && auto.Enable {
			if auto.When == setup.BEFORE {
				auto.Run()
			} else {
				defer auto.Run()
			}
		}
	}

	return nil
}

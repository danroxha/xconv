package initialize

import (
	"fmt"

	"github.com/dannrocha/czen/src/cli"
	"github.com/dannrocha/czen/src/setup"
)


func Execute() error {
	scrip := setup.Script{}
	scrip.LoadScript()

	for _, auto := range scrip.Automation {
		if auto.Bind == cmd.INIT && auto.Enable {
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

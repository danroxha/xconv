package example

import (
	"fmt"

	"github.com/dannrocha/czen/src/cli"
	"github.com/dannrocha/czen/src/setup"
)


func Execute(args ...string) error {
	scrip := setup.Script{}
	scrip.LoadScript()

	for _, auto := range scrip.Automation {
		if auto.Bind == cmd.EXAMPLE && auto.Enable {
			if auto.When == setup.BEFORE {
				auto.Run()
			} else {
				defer auto.Run()
			}
		}
	}

	role := setup.Role{}

	errConf := role.LoadRole()

	if errConf != nil {
		panic(errConf)
	}

	profile, errProf := role.FindCurrentProfileEnable()

	if errProf != nil {
		panic(errProf)
	}

	fmt.Println(profile.Example)

	return nil
}
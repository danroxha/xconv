package cli

import (
	"fmt"

	"github.com/dannrocha/czen/src/setup"
	"github.com/urfave/cli/v2"
)

func Example(c *cli.Context) error {

	scrip := setup.Script{}
	scrip.LoadScript()

	for _, auto := range scrip.Automation {
		if auto.Bind == EXAMPLE && auto.Enable {
			if auto.When == setup.BEFORE {
				auto.Run()
			} else {
				defer auto.Run()
			}
		}
	}

	conf := setup.Configuration{}

	errConf := conf.LoadConfigurationFile()

	if errConf != nil {
		panic(errConf)
	}

	profile, errProf := conf.FindCurrentProfileEnable()

	if errProf != nil {
		panic(errProf)
	}

	fmt.Println(profile.Example)

	return nil
}

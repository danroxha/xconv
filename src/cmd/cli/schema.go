package cli

import (
	"fmt"

	"github.com/dannrocha/czen/src/setup"
	"github.com/urfave/cli/v2"
)

func Schema(c *cli.Context) error {

	script := setup.NewScript()

	
	for _, auto := range script.Automation {
		if auto.Bind == SCHEMA && auto.Enable {
			if auto.When == setup.BEFORE {
				auto.Run()
			} else {
				defer auto.Run()
			}
		}
	}
		
	rule := setup.NewRule()

	profile, errProf := rule.FindCurrentProfileEnable()

	if errProf != nil {
		panic(errProf)
	}

	fmt.Println(profile.Schema)

	return nil
}

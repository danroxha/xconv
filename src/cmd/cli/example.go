package cli

import (
	"fmt"

	"github.com/dannrocha/czen/src/config"
	"github.com/urfave/cli/v2"
)

func Example(c *cli.Context) error {
	conf := config.Configuration{}

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
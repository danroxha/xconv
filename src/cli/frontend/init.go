package frontend

import (
	"github.com/dannrocha/czen/src/cli/command/initialize"
	"github.com/urfave/cli/v2"
)

func Init(c *cli.Context) error {
	return initialize.Execute()
}

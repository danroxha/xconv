package frontend

import (
	"github.com/dannrocha/czen/src/cli/command/version"
	"github.com/urfave/cli/v2"
)

func Version(c *cli.Context) error {
	return version.Execute()
}

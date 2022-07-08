package frontend

import (
	"github.com/dannrocha/czen/src/cli/command/bump"
	"github.com/urfave/cli/v2"
)

func Bump(c *cli.Context) error {
	return bump.Execute()
}

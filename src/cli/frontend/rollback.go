package frontend

import (
	"github.com/dannrocha/czen/src/cli/command/rollback"
	"github.com/urfave/cli/v2"
)

func Rollback(c *cli.Context) error {
	return rollback.Execute()
}

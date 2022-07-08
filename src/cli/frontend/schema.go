package frontend

import (
	"github.com/dannrocha/czen/src/cli/command/schema"
	"github.com/urfave/cli/v2"
)

func Schema(c *cli.Context) error {
	return schema.Execute()
}

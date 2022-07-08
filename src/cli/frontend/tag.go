package frontend

import (
	"github.com/dannrocha/czen/src/cli/command/tag"
	"github.com/urfave/cli/v2"
)

func Tag(c *cli.Context) error {
	return tag.Execute()
}

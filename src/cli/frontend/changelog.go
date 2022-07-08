package frontend

import (
	"github.com/dannrocha/czen/src/cli/command/changelog"
	"github.com/urfave/cli/v2"
)

func Changelog(c *cli.Context) error {
	return changelog.Execute()
}

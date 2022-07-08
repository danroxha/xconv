package frontend

import (
	"github.com/dannrocha/czen/src/cli/command/commit"

	"github.com/urfave/cli/v2"
)

func Commit(c *cli.Context) error {

	return commit.Execute()
}

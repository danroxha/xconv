package frontend

import (
	"github.com/dannrocha/czen/src/cli/command/example"
	"github.com/urfave/cli/v2"
)

func Example(c *cli.Context) error {
	return example.Execute()
}

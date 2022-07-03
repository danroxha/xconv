package application

import (
	"log"
	"os"
	"time"

	"github.com/dannrocha/czen/src/cmd/cli"
	CLI "github.com/urfave/cli/v2"
)

func Run() {
	CLI.VersionFlag = &CLI.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print only the version",
	}

	app := &CLI.App{
		Name:     "czen",
		Compiled: time.Now(),
		Authors: []*CLI.Author{
			{
				Name:  "Rocha da Silva, Daniel",
				Email: "rochadaniel@acad.ifma.edu.br",
			},
		},
		Copyright: "(c) 2022 MIT",
		HelpName:  "contrive",
		Usage:     "Commit ZEN is a cli tool to generate conventional commits.",
		UsageText: "czen [-h] {init,commit,example,info,tag,schema,bump,changelog,version}",
		ArgsUsage: "[args and such]",
		HideHelp:  false,
		Commands: []*CLI.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Action:  cli.Init,
				Usage:   "init commitizen configuration",
			},
			{
				Name:    "commit",
				Aliases: []string{"c"},
				Action:  cli.Commit,
				Usage:   "create new commit",
			},
			{
				Name:    "changelog",
				Aliases: []string{"ch"},
				Action:  cli.Changelog,
				Usage:   "generate changelog (note that it will overwrite existing file)",
			},
			{
				Name:    "bump",
				Aliases: []string{"b"},
				Action:  cli.Bump,
				Usage:   "bump semantic version based on the git log",
			},
			{
				Name:    "rollback",
				Aliases: []string{"r"},
				Action:  cli.Rollback,
				Usage:   "revert commit to a specific tag",
			},
			{
				Name:    "tag",
				Aliases: []string{"t"},
				Action:  cli.Tag,
				Usage:   "show tags",
			},
			{
				Name:    "schema",
				Aliases: []string{"s"},
				Action:  cli.Schema,
				Usage:   "show commit schema",
			},
			{
				Name:    "example",
				Aliases: []string{"e"},
				Action:  cli.Example,
				Usage:   "show commit example",
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Action:  cli.Schema,
				Usage:   "get the version of the installed czen or the current project",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

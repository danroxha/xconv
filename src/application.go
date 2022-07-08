package application

import (
	"log"
	"os"
	"time"

	"github.com/dannrocha/czen/src/cli/command"
	"github.com/dannrocha/czen/src/cli/frontend"
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
				Name:    command.INIT,
				Aliases: []string{"i"},
				Action:  frontend.Init,
				Usage:   "init czen configuration",
			},
			{
				Name:    command.COMMIT,
				Aliases: []string{"c"},
				Action:  frontend.Commit,
				Usage:   "create new commit",
			},
			{
				Name:    command.CHANGELOG,
				Aliases: []string{"ch"},
				Action:  frontend.Changelog,
				Usage:   "generate changelog (note that it will overwrite existing file)",
			},
			{
				Name:    command.BUMP,
				Aliases: []string{"b"},
				Action:  frontend.Bump,
				Usage:   "bump semantic version based on the git log",
			},
			{
				Name:    command.ROLLBACK,
				Aliases: []string{"r"},
				Action:  frontend.Rollback,
				Usage:   "revert commit to a specific tag",
			},
			{
				Name:    command.TAG,
				Aliases: []string{"t"},
				Action:  frontend.Tag,
				Usage:   "show tags",
			},
			{
				Name:    command.SCHEMA,
				Aliases: []string{"s"},
				Action:  frontend.Schema,
				Usage:   "show commit schema",
			},
			{
				Name:    command.EXAMPLE,
				Aliases: []string{"e"},
				Action:  frontend.Example,
				Usage:   "show commit example",
			},
			{
				Name:    command.VERSION,
				Aliases: []string{"v"},
				Action:  frontend.Version,
				Usage:   "get the version of the installed czen or the current project",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/urfave/cli/v2"

	"github.com/postfreely/postfreely"
)

var (
	cmdTemplates cli.Command = cli.Command{
		Name:  "templates",
		Usage: "template management tools",
		Subcommands: []*cli.Command{
			&cmdTemplatesGenerate,
		},
	}

	cmdTemplatesGenerate cli.Command = cli.Command{
		Name:    "generate",
		Aliases: []string{"gen"},
		Usage:   "Generate an initial set of templates",
		Action:  genTemplatesAction,
	}
)

func genTemplatesAction(c *cli.Context) error {
	return postfreely.UnpackTemplates()
}

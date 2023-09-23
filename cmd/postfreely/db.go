/*
 * Copyright © 2020-2021 Musing Studio LLC.
 *
 * This file is part of WriteFreely.
 *
 * WriteFreely is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License, included
 * in the LICENSE file in this source code package.
 */

package main

import (
	"github.com/postfreely/postfreely"
	"github.com/urfave/cli/v2"
)

var (
	cmdDB cli.Command = cli.Command{
		Name:  "db",
		Usage: "db management tools",
		Subcommands: []*cli.Command{
			&cmdDBInit,
			&cmdDBMigrate,
		},
	}

	cmdDBInit cli.Command = cli.Command{
		Name:   "init",
		Usage:  "Initialize Database",
		Action: initDBAction,
	}

	cmdDBMigrate cli.Command = cli.Command{
		Name:   "migrate",
		Usage:  "Migrate Database",
		Action: migrateDBAction,
	}
)

func initDBAction(c *cli.Context) error {
	app := postfreely.NewApp(c.String("c"))
	return postfreely.CreateSchema(app)
}

func migrateDBAction(c *cli.Context) error {
	app := postfreely.NewApp(c.String("c"))
	return postfreely.Migrate(app)
}

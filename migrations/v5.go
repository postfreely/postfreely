/*
 * Copyright © 2019-2021 Musing Studio LLC.
 *
 * This file is part of WriteFreely.
 *
 * WriteFreely is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License, included
 * in the LICENSE file in this source code package.
 */

package migrations

import (
	"context"
	"database/sql"

	dbase "github.com/writefreely/writefreely/db"
)

func oauthSlack(db *datastore) error {
	dialect := dbase.DialectMySQL
	if db.driverName == dbase.TypeSQLite {
		dialect = dbase.DialectSQLite
	}
	return dbase.RunTransactionWithOptions(context.Background(), db.DB, &sql.TxOptions{}, func(ctx context.Context, tx *sql.Tx) error {
		builders := []dbase.SQLBuilder{
			dialect.
				AlterTable("oauth_client_states").
				AddColumn(dialect.
					Column(
						"provider",
						dbase.ColumnTypeVarChar,
						dbase.OptionalInt{Set: true, Value: 24}).SetDefault("")),
			dialect.
				AlterTable("oauth_client_states").
				AddColumn(dialect.
					Column(
						"client_id",
						dbase.ColumnTypeVarChar,
						dbase.OptionalInt{Set: true, Value: 128}).SetDefault("")),
			dialect.
				AlterTable("oauth_users").
				AddColumn(dialect.
					Column(
						"provider",
						dbase.ColumnTypeVarChar,
						dbase.OptionalInt{Set: true, Value: 24}).SetDefault("")),
			dialect.
				AlterTable("oauth_users").
				AddColumn(dialect.
					Column(
						"client_id",
						dbase.ColumnTypeVarChar,
						dbase.OptionalInt{Set: true, Value: 128}).SetDefault("")),
			dialect.
				AlterTable("oauth_users").
				AddColumn(dialect.
					Column(
						"access_token",
						dbase.ColumnTypeVarChar,
						dbase.OptionalInt{Set: true, Value: 512}).SetDefault("")),
			dialect.CreateUniqueIndex("oauth_users_uk", "oauth_users", "user_id", "provider", "client_id"),
		}

		if dialect != dbase.DialectSQLite {
			// This updates the length of the `remote_user_id` column. It isn't needed for SQLite databases.
			builders = append(builders, dialect.
				AlterTable("oauth_users").
				ChangeColumn("remote_user_id",
					dialect.
						Column(
							"remote_user_id",
							dbase.ColumnTypeVarChar,
							dbase.OptionalInt{Set: true, Value: 128})))
		}

		for _, builder := range builders {
			query, err := builder.ToSQL()
			if err != nil {
				return err
			}
			if _, err := tx.ExecContext(ctx, query); err != nil {
				return err
			}
		}
		return nil
	})
}

/*
 * Copyright © 2020-2021 Musing Studio LLC.
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

  dbase "github.com/postfreely/postfreely/db"
)

func oauthAttach(db *datastore) error {
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
						"attach_user_id",
						dbase.ColumnTypeInteger,
						dbase.OptionalInt{Set: true, Value: 24}).SetNullable(true)),
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

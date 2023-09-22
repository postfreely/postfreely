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

func oauth(db *datastore) error {
	dialect := dbase.DialectMySQL
	if db.driverName == dbase.TypeSQLite {
		dialect = dbase.DialectSQLite
	}
	return dbase.RunTransactionWithOptions(context.Background(), db.DB, &sql.TxOptions{}, func(ctx context.Context, tx *sql.Tx) error {
		createTableUsersOauth, err := dialect.
			Table("oauth_users").
			SetIfNotExists(false).
			Column(dialect.Column("user_id", dbase.ColumnTypeInteger, dbase.UnsetSize)).
			Column(dialect.Column("remote_user_id", dbase.ColumnTypeInteger, dbase.UnsetSize)).
			ToSQL()
		if err != nil {
			return err
		}
		createTableOauthClientState, err := dialect.
			Table("oauth_client_states").
			SetIfNotExists(false).
			Column(dialect.Column("state", dbase.ColumnTypeVarChar, dbase.OptionalInt{Set: true, Value: 255})).
			Column(dialect.Column("used", dbase.ColumnTypeBool, dbase.UnsetSize)).
			Column(dialect.Column("created_at", dbase.ColumnTypeDateTime, dbase.UnsetSize).SetDefaultCurrentTimestamp()).
			UniqueConstraint("state").
			ToSQL()
		if err != nil {
			return err
		}

		for _, table := range []string{createTableUsersOauth, createTableOauthClientState} {
			if _, err := tx.ExecContext(ctx, table); err != nil {
				return err
			}
		}
		return nil
	})
}

/*
 * Copyright © 2020 Musing Studio LLC.
 *
 * This file is part of WriteFreely.
 *
 * WriteFreely is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License, included
 * in the LICENSE file in this source code package.
 */

package migrations

import (
	dbase "github.com/postfreely/postfreely/db"
)

func optimizeDrafts(db *datastore) error {
	t, err := db.Begin()
	if err != nil {
		if t != nil {
			t.Rollback()
		}
		return err
	}

	if db.driverName == dbase.TypeSQLite {
		_, err = t.Exec(`CREATE INDEX key_owner_post_id ON posts (owner_id, id)`)
	} else {
		_, err = t.Exec(`ALTER TABLE posts ADD INDEX(owner_id, id)`)
	}
	if err != nil {
		t.Rollback()
		return err
	}

	err = t.Commit()
	if err != nil {
		t.Rollback()
		return err
	}

	return nil
}

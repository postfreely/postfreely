/*
 * Copyright © 2018-2023 PostFreely authors.
 *
 * This file is part of PostFreely.
 *
 * PostFreely is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License, included
 * in the LICENSE file in this source code package.
 */

package usrlvl

import (
	"github.com/postfreely/postfreely/config"
)

type UserLevelFunc func(cfg *config.Config) UserLevel

// UserLevel represents the required user level for accessing an endpoint
type UserLevel int

const (
	UserLevelNoneType         UserLevel = iota // user or not -- ignored
	UserLevelOptionalType                      // user or not -- object fetched if user
	UserLevelNoneRequiredType                  // non-user (required)
	UserLevelUserType                          // user (required)
)

func UserLevelNone(cfg *config.Config) UserLevel {
	return UserLevelNoneType
}

func UserLevelOptional(cfg *config.Config) UserLevel {
	return UserLevelOptionalType
}

func UserLevelNoneRequired(cfg *config.Config) UserLevel {
	return UserLevelNoneRequiredType
}

func UserLevelUser(cfg *config.Config) UserLevel {
	return UserLevelUserType
}

// UserLevelReader returns the permission level required for any route where
// users can read published content.
func UserLevelReader(cfg *config.Config) UserLevel {
	if cfg.App.Private {
		return UserLevelUserType
	}
	return UserLevelOptionalType
}

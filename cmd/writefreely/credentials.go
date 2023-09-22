/*
 * Copyright © 2018-2023 PostFreely authors.
 *
 * This file is part of PostFreely.
 *
 * PostFreely is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License, included
 * in the LICENSE file in this source code package.
 */

package main

import (
	"errors"
	"strings"
)

var (
	errCredentialInvalidFormat = errors.New("invalid format for passed credentials, must be username:password")
)

func parseCredentials(credentialString string) (string, string, error) {
	creds := strings.Split(credentialString, ":")
	if len(creds) != 2 {
		return "", "", errCredentialInvalidFormat
	}
	return creds[0], creds[1], nil
}

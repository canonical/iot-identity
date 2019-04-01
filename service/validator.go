// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * This file is part of the IoT Identity Service
 * Copyright 2019 Canonical Ltd.
 *
 * This program is free software: you can redistribute it and/or modify it
 * under the terms of the GNU Affero General Public License version 3, as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT
 * ANY WARRANTY; without even the implied warranties of MERCHANTABILITY,
 * SATISFACTORY QUALITY, or FITNESS FOR A PARTICULAR PURPOSE.
 * See the GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package service

import (
	"fmt"
	"strings"
)

func normalize(fieldName string) string {
	theFieldName := strings.TrimSpace(fieldName)
	if len(theFieldName) == 0 {
		theFieldName = "field"
	}
	return theFieldName
}

func validateNotEmpty(fieldName, fieldValue string) error {
	if len(strings.TrimSpace(fieldValue)) == 0 {
		return fmt.Errorf("%v must not be empty", normalize(fieldName))
	}
	return nil
}

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

package datastore

import (
	"testing"

	"github.com/CanonicalLtd/iot-identity/config"
)

func TestDatastore_New(t *testing.T) {
	settings := config.Settings{Driver: "memory"}

	db, err := New(&settings)
	if err != nil {
		t.Errorf("datastore.New() error = %v", err)
	}

	if db == nil {
		t.Errorf("datastore.New() error = memory not created")
	}
}

func TestDatastore_NewNegative(t *testing.T) {
	settings := config.Settings{Driver: "garbage"}

	_, err := New(&settings)
	if err == nil {
		t.Errorf("datastore.New() error = no err for incorrect driver")
	}
}

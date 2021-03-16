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

package main

import (
	"log"

	"github.com/bugraaydogar/iot-identity/service/factory"

	"github.com/bugraaydogar/iot-identity/config"
	"github.com/bugraaydogar/iot-identity/service"
	"github.com/bugraaydogar/iot-identity/web"
)

func main() {
	settings := config.ParseArgs()

	// Open the connection to the database
	db, err := factory.CreateDataStore(settings)
	if err != nil {
		log.Fatalf("Error accessing data store: %v", settings.Driver)
	}

	srv := service.NewIdentityService(settings, db)

	// Start the web service
	w := web.NewIdentityService(settings, srv)
	log.Fatal(w.Run())
}

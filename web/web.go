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

package web

import (
	"fmt"
	"net/http"

	"github.com/CanonicalLtd/iot-identity/config"
	"github.com/CanonicalLtd/iot-identity/service"
	"github.com/gorilla/mux"
)

// Web is the interface for the web API
type Web interface {
	Run() error
	Router() *mux.Router
	RegisterOrganization(w http.ResponseWriter, r *http.Request)
	RegisterDevice(w http.ResponseWriter, r *http.Request)
	OrganizationList(w http.ResponseWriter, r *http.Request)
	DeviceList(w http.ResponseWriter, r *http.Request)

	EnrollDevice(w http.ResponseWriter, r *http.Request)
}

// IdentityService is the implementation of the web API
type IdentityService struct {
	Settings *config.Settings
	Identity service.Identity
}

// NewIdentityService returns a new web controller
func NewIdentityService(settings *config.Settings, id service.Identity) *IdentityService {
	return &IdentityService{
		Settings: settings,
		Identity: id,
	}
}

// Run starts the web service
func (wb IdentityService) Run() error {
	fmt.Printf("Starting service on port :%s\n", wb.Settings.Port)
	return http.ListenAndServe(":"+wb.Settings.Port, wb.Router())
}

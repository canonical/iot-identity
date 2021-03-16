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
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/bugraaydogar/iot-identity/service"
)

// RegisterOrganization registers a new organization with the identity service
func (wb IdentityService) RegisterOrganization(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON body
	req, err := decodeOrganizationRequest(w, r)
	if err != nil {
		return
	}

	id, err := wb.Identity.RegisterOrganization(req)
	if err != nil {
		log.Println("Error registering organization:", err)
		formatStandardResponse("RegOrg", err.Error(), w)
		return
	}
	formatRegisterResponse(id, w)
}

// OrganizationList fetches organizations
func (wb IdentityService) OrganizationList(w http.ResponseWriter, r *http.Request) {
	orgs, err := wb.Identity.OrganizationList()
	if err != nil {
		log.Println("Error listing organizations:", err)
		formatStandardResponse("OrgList", err.Error(), w)
		return
	}
	formatOrganizationsResponse(orgs, w)
}

func decodeOrganizationRequest(w http.ResponseWriter, r *http.Request) (*service.RegisterOrganizationRequest, error) { // Decode the REST request
	defer r.Body.Close()

	// Decode the JSON body
	org := service.RegisterOrganizationRequest{}
	err := json.NewDecoder(r.Body).Decode(&org)
	switch {
	// Check we have some data
	case err == io.EOF:
		formatStandardResponse("NoData", "No data supplied.", w)
		log.Println("No data supplied.")
		// Check for parsing errors
	case err != nil:
		formatStandardResponse("BadData", err.Error(), w)
		log.Println(err)
	}
	return &org, err
}

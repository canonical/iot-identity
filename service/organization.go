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
	"github.com/canonical/iot-identity/datastore"
	"github.com/canonical/iot-identity/domain"
	"github.com/canonical/iot-identity/service/cert"
)

// RegisterOrganization registers a new organization with the service
func (id IdentityService) RegisterOrganization(req *RegisterOrganizationRequest) (string, error) {
	// Validate fields
	if err := validateNotEmpty("organization name", req.Name); err != nil {
		return "", err
	}

	// Check that the organization isn't registered i.e. no error with the 'get'
	if _, err := id.DB.OrganizationGetByName(req.Name); err == nil {
		return "", fmt.Errorf("the organization '%s' has already been registered", req.Name)
	}

	// Create server certificate for the organization
	serverPEM, serverCA, err := cert.CreateOrganizationCert(id.Settings.RootCertsDir, req.Name)
	if err != nil {
		return "", err
	}

	// Create registration
	o := datastore.OrganizationNewRequest{
		Name:        req.Name,
		CountryName: req.CountryName,
		ServerKey:   serverPEM,
		ServerCert:  serverCA,
	}

	// Register the organization
	return id.DB.OrganizationNew(o)
}

// OrganizationList fetches the existing organizations
func (id IdentityService) OrganizationList() ([]domain.Organization, error) {
	return id.DB.OrganizationList()
}

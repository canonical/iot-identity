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

package postgres

import (
	"github.com/CanonicalLtd/iot-identity/datastore"
	"github.com/CanonicalLtd/iot-identity/domain"
	"log"
)

// createOrganizationTable creates the database table for organizations with its indexes
func (db *Store) createOrganizationTable() error {
	_, err := db.Exec(createOrganizationTableSQL)
	return err
}

// OrganizationNew creates a new organization
func (db *Store) OrganizationNew(org datastore.OrganizationNewRequest) (string, error) {
	var id int64
	var orgID = datastore.GenerateID()
	err := db.QueryRow(createOrganizationSQL, orgID, org.Name, org.CountryName, org.ServerCert, org.ServerKey).Scan(&id)
	if err != nil {
		log.Printf("Error creating organization: %v\n", err)
	}

	return orgID, err
}

// OrganizationGet fetches an organization by ID
func (db *Store) OrganizationGet(orgID string) (*domain.Organization, error) {
	var id int64
	var countryName string
	org := domain.Organization{}

	err := db.QueryRow(getOrganizationSQL, orgID).Scan(&id, &org.ID, &org.Name, &countryName, &org.RootCert, &org.RootKey)
	if err != nil {
		log.Printf("Error retrieving organization %v: %v\n", orgID, err)
	}
	return &org, err
}

// OrganizationGetByName fetches an organization by name
func (db *Store) OrganizationGetByName(name string) (*domain.Organization, error) {
	var id int64
	var countryName string
	org := domain.Organization{}

	err := db.QueryRow(getOrganizationByNameSQL, name).Scan(&id, &org.ID, &org.Name, &countryName, &org.RootCert, &org.RootKey)
	if err != nil {
		log.Printf("Error retrieving organization `%v`: %v\n", name, err)
	}
	return &org, err
}

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
	"database/sql"

	"github.com/CanonicalLtd/iot-identity/config"
	"github.com/CanonicalLtd/iot-identity/datastore/common"
	"github.com/CanonicalLtd/iot-identity/domain"
	_ "github.com/lib/pq" // postgresql driver
)

// Store implements a PostgreSQL data store
type Store struct {
	Settings *config.Settings
	*sql.DB
}

// NewStore creates a new memory store
func NewStore(settings *config.Settings) *Store {
	return &Store{
		Settings: settings,
	}
}

func (pg *Store) createTables() error {
	_, err := pg.Exec(createOrganizationTableSQL)
	return err
}

// OrganizationNew creates a new organization
func (pg *Store) OrganizationNew(organization common.OrganizationNewRequest) (string, error) {
	panic("implement me")
}

// OrganizationGet fetches an organization by ID
func (pg *Store) OrganizationGet(id string) (*domain.Organization, error) {
	panic("implement me")
}

// OrganizationGetByName fetches an organization by name
func (pg *Store) OrganizationGetByName(name string) (*domain.Organization, error) {
	panic("implement me")
}

// DeviceNew creates a new device registration
func (pg *Store) DeviceNew(device common.DeviceNewRequest) (string, error) {
	panic("implement me")
}

// DeviceGet fetches a device registration
func (pg *Store) DeviceGet(brand, model, serial string) (*domain.Enrollment, error) {
	panic("implement me")
}

// DeviceEnroll enrols a device with the IoT service
func (pg *Store) DeviceEnroll(device common.DeviceEnrollRequest) (*domain.Enrollment, error) {
	panic("implement me")
}

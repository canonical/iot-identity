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
	"errors"
	"github.com/CanonicalLtd/iot-identity/config"
	"github.com/CanonicalLtd/iot-identity/datastore/common"
	"github.com/CanonicalLtd/iot-identity/datastore/memory"
	"github.com/CanonicalLtd/iot-identity/domain"
)

// DataStore is the interfaces for the data repository
type DataStore interface {
	OrganizationNew(organization common.OrganizationNewRequest) (string, error)
	OrganizationGet(id string) (*domain.Organization, error)
	OrganizationGetByName(name string) (*domain.Organization, error)
	DeviceNew(device common.DeviceNewRequest) (string, error)
	DeviceGet(brand, model, serial string) (*domain.Enrollment, error)
	DeviceEnroll(device common.DeviceEnrollRequest) (*domain.Enrollment, error)
}

// Factory method to create data store based on driver selected in settings.
func New(settings *config.Settings) (DataStore, error) {
	if settings.Driver == "memory" {
		db := memory.NewStore()
		return db, nil
	}

	return nil, errors.New("Unknown DataStore driver supplied")
}

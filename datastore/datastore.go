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
	"github.com/CanonicalLtd/iot-identity/domain"
	"github.com/segmentio/ksuid"
)

// DataStore is the interfaces for the data repository
type DataStore interface {
	OrganizationNew(organization OrganizationNewRequest) (string, error)
	OrganizationGet(id string) (*domain.Organization, error)
	OrganizationGetByName(name string) (*domain.Organization, error)
	OrganizationList() ([]domain.Organization, error)

	DeviceNew(device DeviceNewRequest) (string, error)
	DeviceGet(brand, model, serial string) (*domain.Enrollment, error)
	DeviceEnroll(device DeviceEnrollRequest) (*domain.Enrollment, error)
}

// OrganizationNewRequest is the request to create a new organization
type OrganizationNewRequest struct {
	Name        string
	CountryName string
	ServerKey   []byte
	ServerCert  []byte
}

// DeviceNewRequest is the request to create a new device
type DeviceNewRequest struct {
	ID             string
	OrganizationID string
	Brand          string
	Model          string
	SerialNumber   string
	Credentials    domain.Credentials
}

// DeviceEnrollRequest is the request to enroll a device.
// The details come from the model and serial assertion
type DeviceEnrollRequest struct {
	Brand        string
	Model        string
	SerialNumber string
	DeviceKey    string
	StoreID      string
}

// GenerateID generates a unique ID
func GenerateID() string {
	id := ksuid.New()
	return id.String()
}

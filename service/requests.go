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

import "github.com/snapcore/snapd/asserts"

// RegisterOrganizationRequest is the request to create a new organization
type RegisterOrganizationRequest struct {
	Name        string `json:"name"`
	CountryName string `json:"country"`
}

// RegisterDeviceRequest is the request to create a new device
type RegisterDeviceRequest struct {
	OrganizationID string `json:"orgid"`
	Brand          string `json:"brand"`
	Model          string `json:"model"`
	SerialNumber   string `json:"serial"`
}

// EnrollDeviceRequest is the request to enroll a device via assertions
type EnrollDeviceRequest struct {
	Model  asserts.Assertion
	Serial asserts.Assertion
}

// DeviceUpdateRequest holds request to update a device registration
type DeviceUpdateRequest struct {
	Status int `json:"status"`
}

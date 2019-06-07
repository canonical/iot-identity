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
	"github.com/CanonicalLtd/iot-identity/datastore"
	"github.com/CanonicalLtd/iot-identity/domain"
	"github.com/CanonicalLtd/iot-identity/service/cert"
)

// DeviceList fetches the registered devices
func (id IdentityService) DeviceList(orgID string) ([]domain.Enrollment, error) {
	return id.DB.DeviceList(orgID)
}

// RegisterDevice registers a new device with the service
func (id IdentityService) RegisterDevice(req *RegisterDeviceRequest) (string, error) {
	// Validate fields
	for k, v := range map[string]string{
		"organization ID": req.OrganizationID,
		"brand":           req.Brand,
		"model name":      req.Model,
		"serial number":   req.SerialNumber,
	} {
		if err := validateNotEmpty(k, v); err != nil {
			return "", err
		}
	}

	// Check that the organization exists
	org, err := id.DB.OrganizationGet(req.OrganizationID)
	if err != nil {
		return "", err
	}

	// Check that the device has not been registered
	if _, err := id.DB.DeviceGet(req.Brand, req.Model, req.SerialNumber); err == nil {
		return "", fmt.Errorf("the device `%s/%s/%s` is already registered", req.Brand, req.Model, req.SerialNumber)
	}

	// Create a signed certificate
	deviceID := datastore.GenerateID()
	keyPEM, certPEM, err := cert.CreateClientCert(org, id.Settings.RootCertsDir, deviceID)
	if err != nil {
		return "", err
	}

	// Create registration
	d := datastore.DeviceNewRequest{
		ID:             deviceID,
		OrganizationID: req.OrganizationID,
		Brand:          req.Brand,
		Model:          req.Model,
		SerialNumber:   req.SerialNumber,
		Credentials: domain.Credentials{
			PrivateKey:  keyPEM,
			Certificate: certPEM,
			MQTTURL:     id.Settings.MQTTUrl, // Using a default URL for all devices
			MQTTPort:    id.Settings.MQTTPort,
		},
	}
	return id.DB.DeviceNew(d)
}

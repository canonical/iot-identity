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

// DeviceList fetches the registered devices
func (id IdentityService) DeviceList(orgID string) ([]domain.Enrollment, error) {
	return id.DB.DeviceList(orgID)
}

// DeviceGet fetches a device registration
func (id IdentityService) DeviceGet(orgID, deviceID string) (*domain.Enrollment, error) {
	return id.DB.DeviceGetByID(deviceID)
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
		DeviceData: req.DeviceData,
	}
	return id.DB.DeviceNew(d)
}

// DeviceUpdate updates an existing device with the service
// Status changes are limited, depending on whether the device has enrolled with the service. If it has, then it
// already has credentials.
// If a device has not enrolled:
// - Waiting => Disabled
// - Disabled => Waiting
// If a device has enrolled:
// - Enrolled => Disabled (TODO: needs to trigger the removal of credentials from MQTT broker or device or both)
// - Enrolled => Waiting
func (id IdentityService) DeviceUpdate(orgID, deviceID string, req *DeviceUpdateRequest) error {
	// Get the device and check the current status
	device, err := id.DB.DeviceGetByID(deviceID)
	if err != nil {
		return err
	}

	// Update the device data, if it has changed
	if device.DeviceData != req.DeviceData {
		if err := id.DB.DeviceUpdate(device.ID, device.Status, req.DeviceData); err != nil {
			return err
		}
	}

	if req.Status == int(domain.StatusEnrolled) {
		return fmt.Errorf("cannot change a device status to enrolled. The device itself needs to connect for this")
	}

	switch device.Status {
	case domain.StatusWaiting:
		if req.Status == int(domain.StatusWaiting) {
			// No change required
			return nil
		}
		device.Status = domain.StatusDisabled
	case domain.StatusDisabled:
		if req.Status == int(domain.StatusDisabled) {
			// No change required
			return nil
		}
		device.Status = domain.StatusWaiting
	case domain.StatusEnrolled:
		if req.Status == int(domain.StatusDisabled) {
			// TODO: trigger the removal of credentials from MQTT broker or device or both
			device.Status = domain.StatusDisabled
		} else {
			device.Status = domain.StatusWaiting
		}
	}

	return id.DB.DeviceUpdate(device.ID, device.Status, req.DeviceData)
}

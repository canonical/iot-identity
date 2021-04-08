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
	"log"

	"github.com/canonical/iot-identity/config"
	"github.com/canonical/iot-identity/datastore"
	"github.com/canonical/iot-identity/domain"
	"github.com/snapcore/snapd/asserts"
)

// Identity interface for the service
type Identity interface {
	RegisterOrganization(req *RegisterOrganizationRequest) (string, error)
	RegisterDevice(req *RegisterDeviceRequest) (string, error)
	OrganizationList() ([]domain.Organization, error)
	DeviceList(orgID string) ([]domain.Enrollment, error)
	DeviceGet(orgID, deviceID string) (*domain.Enrollment, error)
	DeviceUpdate(orgID, deviceID string, req *DeviceUpdateRequest) error

	EnrollDevice(req *EnrollDeviceRequest) (*domain.Enrollment, error)
}

// IdentityService implementation of the identity use cases
type IdentityService struct {
	Settings *config.Settings
	DB       datastore.DataStore
}

// NewIdentityService creates an implementation of the identity use cases
func NewIdentityService(settings *config.Settings, db datastore.DataStore) *IdentityService {
	return &IdentityService{
		Settings: settings,
		DB:       db,
	}
}

// EnrollDevice connects an IoT device with the service
func (id IdentityService) EnrollDevice(req *EnrollDeviceRequest) (*domain.Enrollment, error) {
	// Validate fields
	if req.Model.Type().Name != asserts.ModelType.Name {
		return nil, fmt.Errorf("the model assertion is an unexpected type")
	}
	if req.Serial.Type().Name != asserts.SerialType.Name {
		return nil, fmt.Errorf("the serial assertion is an unexpected type")
	}

	if req.Model.Header("brand-id") != req.Serial.Header("brand-id") {
		return nil, fmt.Errorf("the brand-id of the model and serial assertion do not match")
	}
	if req.Model.Header("model") != req.Serial.Header("model") {
		return nil, fmt.Errorf("the model name of the model and serial assertion do not match")
	}

	// Create the enrollment request
	enroll := datastore.DeviceEnrollRequest{
		Brand:        req.Model.Header("brand-id").(string),
		Model:        req.Model.Header("model").(string),
		SerialNumber: req.Serial.Header("serial").(string),
		DeviceKey:    req.Serial.Header("device-key").(string),
	}
	if req.Model.Header("store") != nil {
		enroll.StoreID = req.Model.Header("store").(string)
	}

	return id.enroll(&enroll)
}

// Enroll connects an IoT device with the service
func (id IdentityService) enroll(enroll *datastore.DeviceEnrollRequest) (*domain.Enrollment, error) {
	// Get the registration for the device
	dev, err := id.DB.DeviceGet(enroll.Brand, enroll.Model, enroll.SerialNumber)
	if err != nil {
		log.Println("Cannot find registration:", err)
		return nil, err
	}

	// Check that the device is not already enrolled
	switch dev.Status {
	case domain.StatusWaiting:
		break
	case domain.StatusEnrolled:
		return nil, fmt.Errorf("the device `%s/%s/%s` is already enrolled", enroll.Brand, enroll.Model, enroll.SerialNumber)
	case domain.StatusDisabled:
		return nil, fmt.Errorf("the device registration for `%s/%s/%s` is disabled", enroll.Brand, enroll.Model, enroll.SerialNumber)
	default:
		return nil, fmt.Errorf("the device registration for `%s/%s/%s` is invalid", enroll.Brand, enroll.Model, enroll.SerialNumber)
	}

	// Enroll the device
	en, err := id.DB.DeviceEnroll(*enroll)
	if err != nil {
		return nil, err
	}

	// TODO: Register the device in the MQTT broker
	// (Best to do this out-of-band by submitting a message to a queue for processing)

	// Return the MQTT credentials to the device (which includes the unique ID of the device)
	return en, nil
}

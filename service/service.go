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

	"github.com/CanonicalLtd/iot-identity/config"
	"github.com/CanonicalLtd/iot-identity/datastore"
	"github.com/CanonicalLtd/iot-identity/domain"
	"github.com/CanonicalLtd/iot-identity/service/cert"
	"github.com/snapcore/snapd/asserts"
)

// Identity interface for the service
type Identity interface {
	RegisterOrganization(req *RegisterOrganizationRequest) (string, error)
	RegisterDevice(req *RegisterDeviceRequest) (string, error)
	OrganizationList() ([]domain.Organization, error)

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

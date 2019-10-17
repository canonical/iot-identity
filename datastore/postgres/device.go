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
	"fmt"
	"github.com/CanonicalLtd/iot-identity/datastore"
	"github.com/CanonicalLtd/iot-identity/domain"
	"log"
)

// createDeviceTable creates the database table for devices with its indexes
func (db *Store) createDeviceTable() error {
	_, err := db.Exec(createDeviceTableSQL)
	if err != nil {
		return err
	}

	_, err = db.Exec(createDeviceIDIndexSQL)
	if err != nil {
		return err
	}

	_, err = db.Exec(createDeviceBMSIndexSQL)
	if err != nil {
		return err
	}

	// The alter table calls may fail if the field already exists
	_, _ = db.Exec(alterDeviceAddDeviceData)
	return nil
}

// DeviceNew creates a new device registration
func (db *Store) DeviceNew(d datastore.DeviceNewRequest) (string, error) {
	var id int64
	var deviceID = datastore.GenerateID()

	err := db.QueryRow(createDeviceSQL, deviceID, d.OrganizationID, d.Brand, d.Model, d.SerialNumber, d.Credentials.PrivateKey, d.Credentials.Certificate, d.Credentials.MQTTURL, d.Credentials.MQTTPort, d.DeviceData).Scan(&id)
	if err != nil {
		log.Printf("Error creating device: %v\n", err)
	}

	return deviceID, err
}

// DeviceGet fetches a device registration
func (db *Store) DeviceGet(brand, model, serial string) (*domain.Enrollment, error) {
	d := domain.Enrollment{
		Device:       domain.Device{},
		Organization: domain.Organization{},
		Credentials:  domain.Credentials{},
	}

	err := db.QueryRow(getDeviceSQL, brand, model, serial).Scan(
		&d.ID, &d.Organization.ID, &d.Device.Brand, &d.Device.Model, &d.Device.SerialNumber,
		&d.Credentials.PrivateKey, &d.Credentials.Certificate, &d.Credentials.MQTTURL, &d.Credentials.MQTTPort,
		&d.Device.StoreID, &d.Device.DeviceKey, &d.Status, &d.DeviceData)
	if err != nil {
		log.Printf("Error retrieving device: %v\n", err)
		return &d, fmt.Errorf("error retrieving device: %v", err)
	}

	// Get the organization details for the device
	org, err := db.OrganizationGet(d.Organization.ID)
	if err != nil {
		log.Printf("Error retrieving device organization: %v\n", err)
		return &d, fmt.Errorf("error retrieving device organization: %v", err)
	}
	d.Organization = *org

	return &d, err
}

// DeviceGetByID fetches a device registration
func (db *Store) DeviceGetByID(deviceID string) (*domain.Enrollment, error) {
	d := domain.Enrollment{
		Device:       domain.Device{},
		Organization: domain.Organization{},
		Credentials:  domain.Credentials{},
	}

	err := db.QueryRow(getDeviceByIDSQL, deviceID).Scan(
		&d.ID, &d.Organization.ID, &d.Device.Brand, &d.Device.Model, &d.Device.SerialNumber,
		&d.Credentials.PrivateKey, &d.Credentials.Certificate, &d.Credentials.MQTTURL, &d.Credentials.MQTTPort,
		&d.Device.StoreID, &d.Device.DeviceKey, &d.Status)
	if err != nil {
		log.Printf("Error retrieving device: %v\n", err)
		return &d, fmt.Errorf("error retrieving device: %v", err)
	}

	// Get the organization details for the device
	org, err := db.OrganizationGet(d.Organization.ID)
	if err != nil {
		log.Printf("Error retrieving device organization: %v\n", err)
		return &d, fmt.Errorf("error retrieving device organization: %v", err)
	}
	d.Organization = *org

	return &d, err
}

// DeviceEnroll enrolls a device with the IoT service
func (db *Store) DeviceEnroll(d datastore.DeviceEnrollRequest) (*domain.Enrollment, error) {
	_, err := db.Exec(enrollDeviceSQL, d.Brand, d.Model, d.SerialNumber, d.StoreID, d.DeviceKey, domain.StatusEnrolled)
	if err != nil {
		log.Printf("Error updating the device: %v\n", err)
	}

	return db.DeviceGet(d.Brand, d.Model, d.SerialNumber)
}

// DeviceList fetches the device registrations for an organization
func (db *Store) DeviceList(orgID string) ([]domain.Enrollment, error) {
	rows, err := db.Query(listDeviceSQL, orgID)
	if err != nil {
		log.Printf("Error retrieving devices: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	devices := []domain.Enrollment{}
	for rows.Next() {
		d := domain.Enrollment{}
		err := rows.Scan(&d.ID, &d.Organization.ID, &d.Device.Brand, &d.Device.Model, &d.Device.SerialNumber,
			&d.Credentials.Certificate, &d.Credentials.MQTTURL, &d.Credentials.MQTTPort,
			&d.Device.StoreID, &d.Device.DeviceKey, &d.Status, &d.DeviceData)
		if err != nil {
			return nil, err
		}
		devices = append(devices, d)
	}

	return devices, nil
}

// DeviceUpdate updates a device registration
func (db *Store) DeviceUpdate(deviceID string, status domain.Status, deviceData string) error {
	_, err := db.Exec(updateDeviceSQL, deviceID, status, deviceData)
	if err != nil {
		log.Printf("Error updating the device: %v\n", err)
	}

	return err
}

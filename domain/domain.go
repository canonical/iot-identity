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

package domain

// Status is a top-level enrollment status classification
type Status int

// Enrollment status classifications
const (
	StatusWaiting Status = iota + 1
	StatusEnrolled
	StatusDisabled
)

// Organization details for an account
type Organization struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	RootCert []byte `json:"rootcert"`
	RootKey  []byte `json:"-"`
}

// Device details
type Device struct {
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	SerialNumber string `json:"serial"`
	StoreID      string `json:"store"`
	DeviceKey    string `json:"deviceKey"`
}

// Credentials for accessing the MQTT broker
type Credentials struct {
	PrivateKey  []byte `json:"privateKey"`
	Certificate []byte `json:"certificate"`
	MQTTURL     string `json:"mqttUrl"`
	MQTTPort    string `json:"mqttPort"`
}

// Enrollment details for a device
type Enrollment struct {
	ID           string       `json:"id"`
	Device       Device       `json:"device"`
	Credentials  Credentials  `json:"credentials"`
	Organization Organization `json:"organization"`
	Status       Status       `json:"status"`
}

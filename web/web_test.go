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
package web

import (
	"encoding/json"
	"fmt"
	"github.com/canonical/iot-identity/datastore/memory"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/canonical/iot-identity/domain"
	"github.com/canonical/iot-identity/service"
)

type mockIdentity struct {
	withErr bool
}

// RegisterOrganization mocks organization registration
func (id *mockIdentity) RegisterOrganization(req *service.RegisterOrganizationRequest) (string, error) {
	if req.Name == "Exists" {
		return "", fmt.Errorf("MOCK register error")
	}
	return "abc", nil
}

// RegisterDevice mocks device registration
func (id *mockIdentity) RegisterDevice(req *service.RegisterDeviceRequest) (string, error) {
	if req.Brand == "exists" {
		return "", fmt.Errorf("MOCK register error")
	}
	return "def", nil
}

// OrganizationList mocks fetching organizations
func (id *mockIdentity) OrganizationList() ([]domain.Organization, error) {
	if id.withErr {
		return nil, fmt.Errorf("MOCK error list")
	}
	db := memory.NewStore()
	return db.OrganizationList()
}

// DeviceList mocks fetching devices
func (id *mockIdentity) DeviceList(orgID string) ([]domain.Enrollment, error) {
	if id.withErr || orgID == "invalid" {
		return nil, fmt.Errorf("MOCK error list")
	}
	db := memory.NewStore()
	return db.DeviceList(orgID)
}

// DeviceGet mocks fetching a device
func (id *mockIdentity) DeviceGet(orgID, deviceID string) (*domain.Enrollment, error) {
	if id.withErr || deviceID == "invalid" {
		return nil, fmt.Errorf("MOCK error get")
	}
	db := memory.NewStore()
	return db.DeviceGetByID(deviceID)
}

// DeviceUpdate mocks update a device
func (id *mockIdentity) DeviceUpdate(orgID, deviceID string, req *service.DeviceUpdateRequest) error {
	if id.withErr || deviceID == "invalid" {
		return fmt.Errorf("MOCK error update")
	}
	db := memory.NewStore()
	var status domain.Status
	switch req.Status {
	case 2:
		status = domain.StatusEnrolled
	case 3:
		status = domain.StatusDisabled
	default:
		status = domain.StatusWaiting
	}

	return db.DeviceUpdate(deviceID, status, req.DeviceData)
}

// EnrollDevice mocks enrolling a device
func (id *mockIdentity) EnrollDevice(req *service.EnrollDeviceRequest) (*domain.Enrollment, error) {
	return &domain.Enrollment{}, nil
}

func sendRequest(method, url string, data io.Reader, srv *IdentityService) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(method, url, data)

	srv.Router().ServeHTTP(w, r)

	return w
}

func parseRegisterResponse(r io.Reader) (RegisterResponse, error) {
	// Parse the response
	result := RegisterResponse{}
	err := json.NewDecoder(r).Decode(&result)
	return result, err
}

func parseEnrollResponse(r io.Reader) (EnrollResponse, error) {
	// Parse the response
	result := EnrollResponse{}
	err := json.NewDecoder(r).Decode(&result)
	return result, err
}

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
	"io"
	"log"
	"net/http"

	"github.com/bugraaydogar/iot-identity/service"
	"github.com/gorilla/mux"
	"github.com/snapcore/snapd/asserts"
)

// DeviceList fetches device registrations
func (wb IdentityService) DeviceList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	devices, err := wb.Identity.DeviceList(vars["orgid"])
	if err != nil {
		log.Println("Error fetching devices:", err)
		formatStandardResponse("DeviceList", err.Error(), w)
		return
	}
	formatDevicesResponse(devices, w)
}

// DeviceGet fetches a device registration
func (wb IdentityService) DeviceGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	en, err := wb.Identity.DeviceGet(vars["orgid"], vars["device"])
	if err != nil {
		log.Printf("Error fetching device `%s`: %v\n", vars["device"], err)
		formatStandardResponse("DeviceGet", err.Error(), w)
		return
	}
	formatEnrollResponse(*en, w)
}

// DeviceUpdate updates a device registration
func (wb IdentityService) DeviceUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	req, err := decodeDeviceUpdateRequest(w, r)
	if err != nil {
		return
	}

	err = wb.Identity.DeviceUpdate(vars["orgid"], vars["device"], req)
	if err != nil {
		log.Printf("Error updating device `%s`: %v\n", vars["device"], err)
		formatStandardResponse("DeviceUpdate", err.Error(), w)
		return
	}
	formatStandardResponse("", "", w)
}

// RegisterDevice registers a new device with the identity service
func (wb IdentityService) RegisterDevice(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON body
	req, err := decodeDeviceRequest(w, r)
	if err != nil {
		return
	}

	id, err := wb.Identity.RegisterDevice(req)
	if err != nil {
		log.Println("Error registering device:", err)
		formatStandardResponse("RegDevice", err.Error(), w)
		return
	}
	formatRegisterResponse(id, w)
}

// EnrollDevice connects an IoT device with the identity service
func (wb IdentityService) EnrollDevice(w http.ResponseWriter, r *http.Request) {
	// Decode the assertions from the request
	assertion1, assertion2, err := decodeEnrollRequest(r)
	if err != nil {
		formatStandardResponse("EnrollDevice", err.Error(), w)
		return
	}
	if assertion1 == nil || assertion2 == nil {
		formatStandardResponse("EnrollDevice", "A model and serial assertion is required", w)
		return
	}

	req := service.EnrollDeviceRequest{}

	if assertion1.Type().Name == asserts.ModelType.Name && assertion2.Type().Name == asserts.SerialType.Name {
		req.Model = assertion1
		req.Serial = assertion2
	} else if assertion1.Type().Name == asserts.SerialType.Name && assertion2.Type().Name == asserts.ModelType.Name {
		req.Model = assertion2
		req.Serial = assertion1
	}
	if req.Model == nil || req.Serial == nil {
		log.Println("A model and serial assertion must be provided")
	}

	en, err := wb.Identity.EnrollDevice(&req)
	if err != nil {
		log.Println("Error enrolling device:", err)
		formatStandardResponse("EnrollDevice", err.Error(), w)
		return
	}

	formatEnrollResponse(*en, w)
}

func decodeDeviceRequest(w http.ResponseWriter, r *http.Request) (*service.RegisterDeviceRequest, error) { // Decode the REST request
	defer r.Body.Close()

	// Decode the JSON body
	dev := service.RegisterDeviceRequest{}
	err := json.NewDecoder(r.Body).Decode(&dev)
	switch {
	// Check we have some data
	case err == io.EOF:
		formatStandardResponse("NoData", "No data supplied.", w)
		log.Println("No data supplied.")
		// Check for parsing errors
	case err != nil:
		formatStandardResponse("BadData", err.Error(), w)
		log.Println(err)
	}
	return &dev, err
}

func decodeEnrollRequest(r *http.Request) (asserts.Assertion, asserts.Assertion, error) {
	// Use snapd assertion module to decode the assertions in the request stream
	dec := asserts.NewDecoder(r.Body)
	assertion1, err := dec.Decode()
	if err == io.EOF {
		return nil, nil, fmt.Errorf("no data supplied")
	}
	if err != nil {
		return nil, nil, err
	}

	// Decode the second assertion
	assertion2, err := dec.Decode()
	if err != nil && err != io.EOF {
		return nil, nil, err
	}

	// Stream must be ended now
	_, err = dec.Decode()
	if err != io.EOF {
		if err == nil {
			return nil, nil, fmt.Errorf("unexpected assertion in the request stream")
		}
		return nil, nil, err
	}

	return assertion1, assertion2, nil
}

func decodeDeviceUpdateRequest(w http.ResponseWriter, r *http.Request) (*service.DeviceUpdateRequest, error) {
	defer r.Body.Close()

	// Decode the JSON body
	dev := service.DeviceUpdateRequest{}
	err := json.NewDecoder(r.Body).Decode(&dev)
	switch {
	// Check we have some data
	case err == io.EOF:
		formatStandardResponse("NoData", "No data supplied.", w)
		log.Println("No data supplied.")
		// Check for parsing errors
	case err != nil:
		formatStandardResponse("BadData", err.Error(), w)
		log.Println(err)
	}
	return &dev, err
}

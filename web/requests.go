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
	"log"
	"net/http"

	"github.com/CanonicalLtd/iot-identity/domain"
)

// JSONHeader is the header for JSON responses
const JSONHeader = "application/json; charset=UTF-8"

// StandardResponse is the JSON response from an API method, indicating success or failure.
type StandardResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// OrganizationsResponse is the JSON response from a organization list API method
type OrganizationsResponse struct {
	StandardResponse
	Organizations []domain.Organization `json:"organizations"`
}

// DevicesResponse is the JSON response from a device list API method
type DevicesResponse struct {
	StandardResponse
	Devices []domain.Enrollment `json:"devices"`
}

// RegisterResponse is the JSON response from a registration API method
type RegisterResponse struct {
	StandardResponse
	ID string `json:"id"`
}

// EnrollResponse is the JSON response from an enrollment API method
type EnrollResponse struct {
	StandardResponse
	Enrollment domain.Enrollment `json:"enrollment"`
}

// formatStandardResponse returns a JSON response from an API method, indicating success or failure
func formatStandardResponse(code, message string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", JSONHeader)
	response := StandardResponse{Code: code, Message: message}

	if len(code) > 0 {
		w.WriteHeader(http.StatusBadRequest)
	}

	// Encode the response as JSON
	encodeResponse(w, response)
}

// formatOrganizationsResponse returns a JSON response from an organizations API method
func formatOrganizationsResponse(orgs []domain.Organization, w http.ResponseWriter) {
	w.Header().Set("Content-Type", JSONHeader)
	response := OrganizationsResponse{StandardResponse{}, orgs}

	// Encode the response as JSON
	encodeResponse(w, response)
}

// formatDevicesResponse returns a JSON response from an organizations API method
func formatDevicesResponse(items []domain.Enrollment, w http.ResponseWriter) {
	w.Header().Set("Content-Type", JSONHeader)
	response := DevicesResponse{StandardResponse{}, items}

	// Encode the response as JSON
	encodeResponse(w, response)
}

// formatRegisterResponse returns a JSON response from a register API method
func formatRegisterResponse(id string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", JSONHeader)
	response := RegisterResponse{StandardResponse{}, id}

	// Encode the response as JSON
	encodeResponse(w, response)
}

func encodeResponse(w http.ResponseWriter, response interface{}) {
	// Encode the response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Error forming the response:", err)
	}
}

// formatEnrollResponse returns a JSON response from a register API method
func formatEnrollResponse(en domain.Enrollment, w http.ResponseWriter) {
	w.Header().Set("Content-Type", JSONHeader)
	response := EnrollResponse{StandardResponse{}, en}

	// Encode the response as JSON
	encodeResponse(w, response)
}

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
	"testing"

	"github.com/CanonicalLtd/iot-identity/config"
	"github.com/CanonicalLtd/iot-identity/datastore/memory"
)

func TestIdentityService_DeviceList(t *testing.T) {
	settings := &config.Settings{RootCertsDir: "../datastore/test_data"}
	db := memory.NewStore()
	type args struct {
		orgID string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"valid", args{"abc"}, 3, false},
		{"invalid", args{"invalid"}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := NewIdentityService(settings, db)
			got, err := id.DeviceList(tt.args.orgID)
			if (err != nil) != tt.wantErr {
				t.Errorf("IdentityService.DeviceList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("IdentityService.DeviceList() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func TestIdentityService_DeviceGet(t *testing.T) {
	settings := &config.Settings{RootCertsDir: "../datastore/test_data"}
	db := memory.NewStore()
	type args struct {
		orgID    string
		deviceID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"valid", args{"abc", "a111"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := NewIdentityService(settings, db)
			got, err := id.DeviceGet(tt.args.orgID, tt.args.deviceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("IdentityService.DeviceGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.ID != tt.args.deviceID {
				t.Errorf("IdentityService.DeviceGet() = %v, want %v", got.ID, tt.args.deviceID)
			}
		})
	}
}

func TestIdentityService_DeviceUpdate(t *testing.T) {
	settings := &config.Settings{RootCertsDir: "../datastore/test_data"}
	type args struct {
		orgID    string
		deviceID string
		req      *DeviceUpdateRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"valid", args{"abc", "a111", &DeviceUpdateRequest{Status: 3, DeviceData: "abc"}}, false},
		{"invalid-device", args{"abc", "invalid", &DeviceUpdateRequest{Status: 3, DeviceData: "abc"}}, true},
		{"invalid-enrolled", args{"abc", "a111", &DeviceUpdateRequest{Status: 2, DeviceData: "abc"}}, true},
		{"valid-waiting-disabled", args{"abc", "c333", &DeviceUpdateRequest{Status: 3, DeviceData: "abc"}}, false},
		{"valid-waiting-unchanged", args{"abc", "c333", &DeviceUpdateRequest{Status: 1, DeviceData: "abc"}}, false},
		{"valid-enrolled-disabled", args{"abc", "b222", &DeviceUpdateRequest{Status: 3, DeviceData: "abc"}}, false},
		{"valid-enrolled-waiting", args{"abc", "b222", &DeviceUpdateRequest{Status: 1, DeviceData: "abc"}}, false},
		{"valid-disabled-unchanged", args{"abc", "a111", &DeviceUpdateRequest{Status: 3, DeviceData: "abc"}}, false},
		{"valid-disabled-waiting", args{"abc", "a111", &DeviceUpdateRequest{Status: 1, DeviceData: "abc"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := memory.NewStore()
			id := NewIdentityService(settings, db)
			if err := id.DeviceUpdate(tt.args.orgID, tt.args.deviceID, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("IdentityService.DeviceUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

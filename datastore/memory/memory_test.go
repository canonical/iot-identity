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

package memory

import (
	"reflect"
	"testing"

	"github.com/CanonicalLtd/iot-identity/datastore"
	"github.com/CanonicalLtd/iot-identity/domain"
)

func TestStore_OrganizationNew(t *testing.T) {
	req1 := datastore.OrganizationNewRequest{
		Name:        "Example Ltd",
		CountryName: "United Kingdom",
		ServerKey:   []byte("-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEA4LzuMogDv9It"),
		ServerCert:  []byte("-----BEGIN CERTIFICATE-----\nMIICYzCCAUsCAQAwHjEcMBoGA1UECgwT"),
	}
	req2 := datastore.OrganizationNewRequest{}
	req3 := datastore.OrganizationNewRequest{
		Name:        "Example Inc",
		CountryName: "United Kingdom",
		ServerKey:   []byte("-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEA4LzuMogDv9It"),
		ServerCert:  []byte("-----BEGIN CERTIFICATE-----\nMIICYzCCAUsCAQAwHjEcMBoGA1UECgwT"),
	}
	type args struct {
		organization datastore.OrganizationNewRequest
	}
	tests := []struct {
		name    string
		args    args
		count   int
		wantErr bool
	}{
		{"valid", args{req1}, 2, false},
		{"invalid", args{req2}, 1, true},
		{"duplicate", args{req3}, 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStore()
			got, err := s.OrganizationNew(tt.args.organization)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.OrganizationNew() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(got) == 0 {
					t.Error("Store.OrganizationNew() = no ID generated")
				}
			}
			if len(s.Orgs) != tt.count {
				t.Errorf("Store.OrganizationNew() count = %v, want %v", len(s.Orgs), tt.count)
			}
		})
	}
}

func TestStore_DeviceNew(t *testing.T) {
	req1 := datastore.DeviceNewRequest{
		OrganizationID: "abc",
		Brand:          "example",
		Model:          "drone-1000",
		SerialNumber:   "b222",
	}
	req2 := datastore.DeviceNewRequest{}
	req3 := datastore.DeviceNewRequest{
		OrganizationID: "abc",
		Brand:          "example",
		Model:          "drone-1000",
		SerialNumber:   "DR1000A111",
	}
	req4 := datastore.DeviceNewRequest{
		OrganizationID: "invalid",
		Brand:          "example",
		Model:          "drone-1000",
		SerialNumber:   "b222",
	}

	type args struct {
		device datastore.DeviceNewRequest
	}
	tests := []struct {
		name    string
		args    args
		count   int
		wantErr bool
	}{
		{"valid", args{req1}, 4, false},
		{"invalid", args{req2}, 3, true},
		{"duplicate", args{req3}, 3, true},
		{"invalid-org", args{req4}, 3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStore()
			got, err := s.DeviceNew(tt.args.device)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.DeviceNew() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(got) == 0 {
					t.Error("Store.DeviceNew() = no ID generated")
				}
			}
			if len(s.Roll) != tt.count {
				t.Errorf("Store.DeviceNew() count = %v, want %v", len(s.Roll), tt.count)
			}
		})
	}
}

func TestStore_DeviceEnroll(t *testing.T) {
	req1 := datastore.DeviceEnrollRequest{
		Brand:        "example",
		Model:        "drone-1000",
		SerialNumber: "DR1000A111",
		StoreID:      "example-store",
		DeviceKey:    "-----BEGIN GPG PUBLIC KEY-----\nMIIEpAIBAAKCAQ",
	}
	req2 := datastore.DeviceEnrollRequest{
		Brand:        "invalid",
		Model:        "drone-1000",
		SerialNumber: "DR1000A111",
		StoreID:      "example-store",
		DeviceKey:    "-----BEGIN GPG PUBLIC KEY-----\nMIIEpAIBAAKCAQ",
	}

	// Reply
	exOrg := domain.Organization{ID: "abc", Name: "Example Inc", RootKey: []byte(RootPEM), RootCert: []byte(CertPEM)}
	dev1 := domain.Device{Brand: "example", Model: "drone-1000", SerialNumber: "DR1000A111", StoreID: "example-store", DeviceKey: "-----BEGIN GPG PUBLIC KEY-----\nMIIEpAIBAAKCAQ"}
	reply1 := &domain.Enrollment{
		ID:           "a111",
		Organization: exOrg,
		Device:       dev1,
		Status:       domain.StatusEnrolled,
	}

	type args struct {
		device datastore.DeviceEnrollRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Enrollment
		wantErr bool
	}{
		{"valid", args{req1}, reply1, false},
		{"invalid", args{req2}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStore()
			got, err := s.DeviceEnroll(tt.args.device)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.DeviceEnroll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.DeviceEnroll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_OrganizationGetByName(t *testing.T) {
	reply1 := domain.Organization{ID: "abc", Name: "Example Inc", RootKey: []byte(RootPEM), RootCert: []byte(CertPEM)}

	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Organization
		wantErr bool
	}{
		{"valid", args{"Example Inc"}, &reply1, false},
		{"not-found", args{"Non-Existent"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mem := NewStore()
			got, err := mem.OrganizationGetByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.OrganizationGetByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.OrganizationGetByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_OrganizationList(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		{"valid", 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mem := NewStore()
			got, err := mem.OrganizationList()
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.OrganizationList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("Store.OrganizationList() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

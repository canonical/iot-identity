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

	"github.com/bugraaydogar/iot-identity/config"
	"github.com/bugraaydogar/iot-identity/datastore/memory"
)

var settings = config.ParseArgs()

func TestIdentityService_OrganizationList(t *testing.T) {
	settings.RootCertsDir = "../datastore/test_data"
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		{"valid", 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := NewIdentityService(settings, memory.NewStore())
			got, err := id.OrganizationList()
			if (err != nil) != tt.wantErr {
				t.Errorf("IdentityService.OrganizationList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("IdentityService.OrganizationList() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

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

package cert

import (
	"testing"
)

func TestCreateOrganizationCert(t *testing.T) {
	type args struct {
		certsPath string
		orgName   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"valid", args{"../../datastore/test_data", "Example PLC"}, false},
		{"invalid-path", args{"invalid", "Example PLC"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := CreateOrganizationCert(tt.args.certsPath, tt.args.orgName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOrganizationCert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil && !tt.wantErr {
				t.Errorf("CreateOrganizationCert() got = %v, want certificate", got)
			}
			if got1 == nil && !tt.wantErr {
				t.Errorf("CreateOrganizationCert() got1 = %v, want certificate", got1)
			}
		})
	}
}

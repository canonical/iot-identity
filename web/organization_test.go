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
	"bytes"
	"testing"

	"github.com/CanonicalLtd/iot-identity/config"
)

func TestIdentityService_RegisterOrganization(t *testing.T) {
	settings := &config.Settings{}
	req1 := []byte(`{"name":"Example AB", "country":"Sweden"}`)
	req2 := []byte(`{"name":"Exists", "country":"Sweden"}`)
	req3 := []byte(``)
	req4 := []byte(`\u000`)

	type args struct {
		req []byte
	}
	tests := []struct {
		name   string
		args   args
		code   int
		result string
	}{
		{"valid", args{req1}, 200, ""},
		{"duplicate", args{req2}, 400, "RegOrg"},
		{"no-data", args{req3}, 400, "NoData"},
		{"bad-data", args{req4}, 400, "BadData"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wb := NewIdentityService(settings, &mockIdentity{})

			w := sendRequest("POST", "/v1/organization", bytes.NewReader(tt.args.req), wb)
			if w.Code != tt.code {
				t.Errorf("Web.RegisterOrganization() got = %v, want %v", w.Code, tt.code)
			}
			resp, err := parseRegisterResponse(w.Body)
			if err != nil {
				t.Errorf("Web.RegisterOrganization() got = %v", err)
			}
			if resp.Code != tt.result {
				t.Errorf("Web.RegisterOrganization() got = %v, want %v", resp.Code, tt.result)
			}
		})
	}
}

func TestIdentityService_OrganizationList(t *testing.T) {
	tests := []struct {
		name    string
		withErr bool
		code    int
		result  string
	}{
		{"valid", false, 200, ""},
		{"invalid", true, 400, "OrgList"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wb := NewIdentityService(settings, &mockIdentity{tt.withErr})

			w := sendRequest("GET", "/v1/organizations", nil, wb)
			if w.Code != tt.code {
				t.Errorf("Web.OrganizationList() got = %v, want %v", w.Code, tt.code)
			}
			resp, err := parseRegisterResponse(w.Body)
			if err != nil {
				t.Errorf("Web.OrganizationList() got = %v", err)
			}
			if resp.Code != tt.result {
				t.Errorf("Web.OrganizationList() got = %v, want %v", resp.Code, tt.result)
			}
		})
	}
}

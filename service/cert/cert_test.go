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

	"github.com/bugraaydogar/iot-identity/datastore/memory"
)

func Test_getCertificateAuthority(t *testing.T) {
	type args struct {
		certsPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"valid", args{"../../datastore/test_data"}, false},
		{"invalid-path", args{"invalid"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1, err := getCertificateAuthority(tt.args.certsPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCertificateAuthority() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got1 == nil {
					t.Errorf("getCertificateAuthority() got1 = %v, want certificate", got1)
				}
			}
		})
	}
}

func Test_parseRootCertificate(t *testing.T) {
	type args struct {
		rootCert []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"valid", args{[]byte(memory.CertPEM)}, false},
		{"invalid", args{[]byte("invalid")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseRootCertificate(tt.args.rootCert)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseRootCertificate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got == nil || got.SerialNumber == nil {
					t.Errorf("parseRootCertificate() = %v, want CA", got)
				}
			}
		})
	}
}

func Test_certToPEM(t *testing.T) {
	type args struct {
		c []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"valid", args{[]byte(memory.CertPEM)}, false},
		{"valid", args{[]byte("invalid")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := certToPEM(tt.args.c)
			if got == nil && !tt.wantErr {
				t.Errorf("certToPEM() = %v, want success", got)
			}
		})
	}
}

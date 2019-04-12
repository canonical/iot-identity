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
	"crypto/x509"
	"reflect"
	"testing"
)

func Test_getCertificateAuthority(t *testing.T) {
	type args struct {
		certsPath string
	}
	tests := []struct {
		name    string
		args    args
		want1   *x509.Certificate
		wantErr bool
	}{
		{"invalid-path", args{"invalid"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1, err := getCertificateAuthority(tt.args.certsPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCertificateAuthority() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getCertificateAuthority() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

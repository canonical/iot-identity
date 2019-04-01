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

package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"default-settings-create"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			{
				got := ParseArgs()
				assert.Equal(t, DefaultPort, got.Port, tt.name)
				assert.Equal(t, DefaultDriver, got.Driver, tt.name)
				assert.Equal(t, DefaultDataSource, got.DataSource, tt.name)
				assert.Equal(t, DefaultMQTTURL, got.MQTTUrl, tt.name)
				assert.Equal(t, DefaultMQTTPort, got.MQTTPort, tt.name)
				assert.True(t, len(got.KeySecret) > 0, "secret not generated")

				_ = os.Remove(keyFilename)
			}
		})
	}

}

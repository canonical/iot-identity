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

	"github.com/snapcore/snapd/asserts"

	"github.com/CanonicalLtd/iot-identity/config"
	"github.com/CanonicalLtd/iot-identity/datastore"
	"github.com/CanonicalLtd/iot-identity/datastore/memory"
)

const model1 = `type: model
authority-id: canonical
series: 16
brand-id: canonical
model: ubuntu-core-18-amd64
architecture: amd64
base: core18
display-name: Ubuntu Core 18 (amd64)
gadget: pc=18
kernel: pc-kernel=18
timestamp: 2018-08-13T09:00:00+00:00
sign-key-sha3-384: 9tydnLa6MTJ-jaQTFUXEwHl1yRx7ZS4K5cyFDhYDcPzhS7uyEkDxdUjg9g08BtNn

AcLBXAQAAQoABgUCW37NBwAKCRDgT5vottzAEut9D/4u9lD3lFWXoHx1VQT+mUCROcFHdXQBY/PJ
NriRiDwBaOjEo5mvHMRJ2UulWvHnwqyMJctJKBP+RCKlrJEPX8eaLP/lmihwIiFfmzm49BLaNwli
si0entond1sVWfiNr7azXoEuAIgYvxmJIvE+GZADDT0/OTFQRcLU69bhNEAQKBnkT0y/HTpuXwlJ
TuwwJtDR0vZuFtwzj6Bdx7W42+vGmuXE7M4Ni6HUySNKYByB5BsrDf3/79p8huXyBtnWp+HBsHtb
fgjzQoBcspj65Gi+crBrJ4jS+nfowRRVXLL1clXJOJLz12za+kN0/FC0PhussiQb5UI7USXJ+RvA
Y8U1vrqG7bG5GYGqe1KB9GbLEm+GBPQZcZI3jRmm9V7tm9OWQzK98/uPwTD73IW7LrDT35WQrIYM
fBfThJcRqpgzwZD/CBx82maLB9tmsRF5Mhcj2H1v7cn8nSkbv7+cCzh25lKv48Vqz1WTgO3HMPWW
0kb6BSoC+YGpstSUslqtpLdY/MfFI0DhshH2Y+h0c9/g4mux/Zb8Gs9V55HGn9mr2KKDmHsU2k+C
maZWcXOxRpverZ2Pi9L4fZxhZ9H+FDcMGiHn2vJFQhI3u+LiK3aUUAov4k3vNRPGSvi1AGhuEtUa
NG54bznx12KgOT3+YiHtfE95WiXUcJUrEXAgfVBVoA==`

const serial1 = `type: serial
authority-id: canonical
brand-id: canonical
model: ubuntu-core-18-amd64
serial: d75f7300-abbf-4c11-bf0a-8b7103038490
device-key:
    AcbBTQRWhcGAARAA05GC1FmdsBVDxd2DbolPLiqnQXDDwW0RScEcuG5ONGMmvolfS4DJxS5ONBq2
    ZdvGYoCzuSE4P/fruKwrfnR+DRn+frA2YAQOagHy2xmSYlXBz1wyDAvKVmJdv7Q2EjGK4K6vgVMn
    v8No+9/fecoIF7oa9kF7EwcnDrN89VGR+jOljGvwJ3QKHh8Tq5szL3ETlhdv4E6GEt4lEjcw3hDM
    rjGezRwM9riypbJp3paNWygff03sC6Q5esZk9U2ijF7tEF7CT5zCZEaLs+OdOQxYL6R4Bw7lp2h2
    xj/0G6pX3AH/VtijIJj/aOn6fBQB9kzGEghjUemHKqfpJ7lEH/TQ0JIMj9z/Tgj5KDPXEgtwgf78
    37TYbDxcfoFJbi4sMoXFoKq2d2b8ufnQ1UlxMiCxr/z3GtraxDhMRx34vxIr1RqhHGt48as0rLjF
    mnsOAxSOhyloVgd9V5jdK7gzCi6aTtNZTMJV5TkGo3HyMEmDmj+TLAmPrENVt2A/EnKEyORz+0o1
    5qtauqdcypOyAQc1aPmbGtqX5adI8tuj6JLxXdcQgCsQp+F5j+NM9TZnNnbwjkWZam1G8seGH+GZ
    QpeT5+5VqhXIkmlk8Mfqgn5br/1D7dfjBrzAumBpOmcOIeCCYrBtlpva4+nnO3Hp6bmkfuYBNXZe
    jJJS3M6FTNApbr0AEQEAAQ==
device-key-sha3-384: xm9bu3yCuJguaB233yCAnXDE9zgOu8V39-2j8c-Rk0R27HjQpruF8ce_vGZDEm-G
timestamp: 2019-01-10T17:40:44.771564Z
sign-key-sha3-384: BWDEoaqyr25nF5SNCvEv2v7QnM9QsfCc0PBMYD_i2NGSQ32EF2d4D0hqUel3m8ul

AcLBUgQAAQoABgUCXDeDnAAAnLMQAG6jJOffkqDrUhbgMP6VBmGr9nTm54fUg+pMYvxVxex6o4vH
thA5qtQE9of1UVAK5qX7qwwl3rsIZ1/ESagW1ME1hyrCcVxcZ63BQrLODj9VX0kp8VmBvgUWGIsw
sS/ZidF4lbsanWyzFefCErgzAncjxGN9cpMUsJPd5ai2c6Iq9+8qvJoT6ubWWg0Nh/Fe+jURKTs8
Sfzfz0vaySoSmuH4cOYShz2tYvVEVvJyaoNt5vLUrG2TKgA5tz1S0mKwhwDbGRwKFL6mQSlJ/L5N
P6UKSpZKfin+/ziH5YV0PoY3pTeTbuoMQWknYqQUBN/rHzd1y6xmY6rcWsZkFN2sPqA57ZgxUW4C
h/3TZDyRUNXSGqiam5lKEx1EUWiWHhZG6TtOG8+pOW+Y+uW8v1c2qKKHIghQHAgZjUzaNyec2Ylw
PfZW5UO8ua37jvSDV4aYcDXLlumD76mCQkXslltXATOnH9ZDMaf7/MRnx7Dwaqu0kuYUCNSWN/kJ
oe5AnCaMg/yTp0EbV9ZlHNeQYGesUkhT9ULXzsUEfhs3S6mQtnC12O1C/F7fsv1x7lSa4WvPzlb7
Azds7xIR91OzXGFMx/PO7ZwflxBRIZw7+iFXEXWzfhzVlrUFDLr8K++g1g563UzY9P86XwGDlS7l
/PVxRaD/Ruiw0ey94zCcn3ROBEs/`

const model2 = `type: model
authority-id: generic
series: 16
brand-id: generic
model: generic-classic
classic: true
timestamp: 2017-07-27T00:00:00.0Z
sign-key-sha3-384: d-JcZF9nD9eBw7bwMnH61x-bklnQOhQud1Is6o_cn2wTj8EYDi9musrIT9z2MdAa

AcLBXAQAAQoABgUCWYuXiAAKCRAdLQyY+/mCiST0D/0XGQauzV2bbTEy6DkrR1jlNbI6x8vfIdS8
KvEWYvzOWNhNlVSfwNOkFjs3uMHgCO6/fCg03wGXTyV9D7ZgrMeUzWrYp6EmXk8/LQSaBnff86XO
4/vYyfyvEYavhF0kQ6QGg8Cqr0EaMyw0x9/zWEO/Ll9fH/8nv9qcQq8N4AbebNvNxtGsCmJuXpSe
2rxl3Dw8XarYBmqgcBQhXxRNpa6/AgaTNBpPOTqgNA8ZtmbZwYLuaFjpZP410aJSs+evSKepy/ce
+zTA7RB3384YQVeZDdTudX2fGtuCnBZBAJ+NYlk0t8VFXxyOhyMSXeylSpNSx4pCqmUZRyaf5SDS
g1XxJet4IP0stZH1SfPOwc9oE81/bJlKsb9QIQKQRewvtUCLfe9a6Vy/CYd2elvcWOmeANVrJK0m
nRaz6VBm09RJTuwUT6vNugXSOCeF7W3WN1RHJuex0zw+nP3eCehxFSr33YrVniaA7zGfjXvS8tKx
AINNQB4g2fpfet4na6lPPMYM41WHIHPCMTz/fJQ6dZBSEg6UUZ/GiQhGEfWPBteK7yd9pQ8qB3fj
ER4UvKnR7hcVI26e3NGNkXP5kp0SFCkV5NQs8rzXzokpB7p/V5Pnqp3Km6wu45cU6UiTZFhR2IMT
l+6AMtrS4gDGHktOhwfmOMWqmhvR/INF+TjaWbsB6g==
`

const model3 = `type: model
authority-id: canonical
series: 16
brand-id: canonical
model: pc-amd64
architecture: amd64
gadget: pc
kernel: pc-kernel
timestamp: 2016-08-31T00:00:00.0Z
sign-key-sha3-384: 9tydnLa6MTJ-jaQTFUXEwHl1yRx7ZS4K5cyFDhYDcPzhS7uyEkDxdUjg9g08BtNn

AcLBXAQAAQoABgUCV9A82wAKCRDgT5vottzAEhq1D/4z66k0JS7sQrD54Ccros3HaAABF+7KwGqV
ggg6Mk+N2QKNxpl7fxHeyB82KUy49v4Kp8cg4icPUfrZb1DyzjgyuJIzZfCp1+LLQ4ShJ0ZW9MLW
p7r/FbITtbmGlCKjVtaSwLYTkZNfae/MTTuTB1nLXH939vdicRPtRQ1MsoQ6v8wUYeE4/F+SUxL9
ekYf4G8sz+vzcO5BK9+1T3Wo/aLHDi0N4EOS3K4ia1BVITZKvyeIUEHOLQJAHKk43dAL0PqMFW+W
IHhDXQoUeiURBfy6zcrRynaIj5tzlhFmJ3pjlmLQLlVCeGJ4yuZ6xb0YIl+oHpYzZrxTad2mEMUY
si4qIyxVNGj7LZCloLRsDFBMh8RS9a8L0/Cq3hA2Q1Ugyw2D5U7J427SVYCDS9rrihNVvMFscou6
vrZHMnAVl/F/TRUDYy29idiiibBQU02D1l4Qu7QnDQQCZygq1n+aeW5ZPwtF/KclkJm0YRUkqbtR
FG2TYLmQ06MPmRuqVRaAdjfhnZ9YtFBDhI+obn99q/OmG2e7d4WNU3JPG1h5arIQGNeR9kVzBER1
iO0V3iYjD0DxOsd2QVOdI/o8HqCRfycTMo/7TydVdWKXKpKdzeezfz/df2LRDCE712NVFhY0hDC6
BvV4mMoqS17K7OMHfDohh0DFfp0yFl9oYfLY55G5HA==`

const ignoreError1 = "the device `canonical/canonical/d75f7300-abbf-4c11-bf0a-8b7103038490` is not registered"

func TestIdentityService_RegisterOrganization(t *testing.T) {
	settings := config.ParseArgs()
	settings.RootCertsDir = "../datastore/test_data"
	db := memory.NewStore()
	req1 := RegisterOrganizationRequest{
		Name:        "Example PLC",
		CountryName: "United Kingdom",
	}
	req2 := RegisterOrganizationRequest{
		Name:        "",
		CountryName: "United Kingdom",
	}
	req3 := RegisterOrganizationRequest{
		Name:        "Example Inc",
		CountryName: "United Kingdom",
	}

	tests := []struct {
		name    string
		req     RegisterOrganizationRequest
		wantErr bool
	}{
		{"valid", req1, false},
		{"invalid", req2, true},
		{"duplicate", req3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := NewIdentityService(settings, db)
			got, err := id.RegisterOrganization(&tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("IdentityService.RegisterOrganization() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(got) == 0 {
					t.Error("Store.OrganizationNew() = no ID generated")
				}
			}
		})
	}
}

func TestIdentityService_RegisterDevice(t *testing.T) {
	settings := &config.Settings{RootCertsDir: "../datastore/test_data"}
	db := memory.NewStore()
	req1 := RegisterDeviceRequest{
		OrganizationID: "abc",
		Brand:          "example",
		Model:          "drone-2000",
		SerialNumber:   "DR2000B2222",
	}
	req2 := RegisterDeviceRequest{
		OrganizationID: "abc",
		Brand:          "",
		Model:          "drone-2000",
		SerialNumber:   "DR2000B2222",
	}
	req3 := RegisterDeviceRequest{
		OrganizationID: "invalid",
		Brand:          "example",
		Model:          "drone-1000",
		SerialNumber:   "DR1000C333",
	}
	req4 := RegisterDeviceRequest{
		OrganizationID: "abc",
		Brand:          "example",
		Model:          "drone-1000",
		SerialNumber:   "DR1000A111",
	}

	type args struct {
		req RegisterDeviceRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"valid", args{req1}, false},
		{"invalid", args{req2}, true},
		{"invalid-org", args{req3}, true},
		{"duplicate-device", args{req4}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := NewIdentityService(settings, db)
			got, err := id.RegisterDevice(&tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("IdentityService.RegisterDevice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(got) == 0 {
					t.Error("IdentityService.RegisterDevice() = no ID generated")
				}
			}
		})
	}
}

func TestIdentityService_Enroll(t *testing.T) {
	settings := &config.Settings{RootCertsDir: "../datastore/test_data"}
	db := memory.NewStore()
	req1 := datastore.DeviceEnrollRequest{
		Brand:        "example",
		Model:        "drone-1000",
		SerialNumber: "DR1000A111",
		StoreID:      "example-store",
		DeviceKey:    "AAAAAAAA",
	}
	req2 := datastore.DeviceEnrollRequest{
		Brand:        "invalid",
		Model:        "drone-1000",
		SerialNumber: "DR1000A111",
		StoreID:      "example-store",
		DeviceKey:    "-----BEGIN GPG PUBLIC KEY-----\nMIIEpAIBAAKCAQ",
	}
	req3 := datastore.DeviceEnrollRequest{
		Brand:        "example",
		Model:        "drone-1000",
		SerialNumber: "DR1000B222",
		StoreID:      "example-store",
		DeviceKey:    "-----BEGIN GPG PUBLIC KEY-----\nMIIEpAIBAAKCAQ",
	}
	req4 := datastore.DeviceEnrollRequest{
		Brand:        "",
		Model:        "drone-1000",
		SerialNumber: "DR1000B222",
		StoreID:      "example-store",
		DeviceKey:    "-----BEGIN GPG PUBLIC KEY-----\nMIIEpAIBAAKCAQ",
	}

	type args struct {
		req datastore.DeviceEnrollRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"valid", args{req1}, false},
		{"invalid", args{req2}, true},
		{"enrolled", args{req3}, true},
		{"empty", args{req4}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := NewIdentityService(settings, db)
			got, err := id.enroll(&tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("IdentityService.Enroll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got == nil {
					t.Error("IdentityService.Enroll() = enrollment is nil")
					return
				}
				if len(got.Device.DeviceKey) == 0 {
					t.Error("IdentityService.Enroll() = device key is not populated")
				}
			}
		})
	}
}

func TestIdentityService_EnrollDevice(t *testing.T) {
	settings := &config.Settings{RootCertsDir: "../datastore/test_data"}
	db := memory.NewStore()

	type args struct {
		model  string
		serial string
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
		ignoreErr string
	}{
		{"valid", args{model1, serial1}, false, ignoreError1},
		{"no-serial", args{model1, model1}, true, ignoreError1},
		{"no-model", args{serial1, serial1}, true, ignoreError1},
		{"mismatch-brand", args{model2, serial1}, true, ignoreError1},
		{"mismatch-model", args{model3, serial1}, true, ignoreError1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := asserts.Decode([]byte(tt.args.model))
			if err != nil {
				t.Errorf("IdentityService.EnrollDevice() model error = %v", err)
			}
			s, err := asserts.Decode([]byte(tt.args.serial))
			if err != nil {
				t.Errorf("IdentityService.EnrollDevice() serial error = %v", err)
			}

			req := &EnrollDeviceRequest{
				Model:  m,
				Serial: s,
			}

			id := NewIdentityService(settings, db)
			got, err := id.EnrollDevice(req)
			if err != nil && !tt.wantErr {
				if err.Error() != tt.ignoreErr {
					t.Errorf("IdentityService.EnrollDevice() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !tt.wantErr {
				if got == nil {
					t.Errorf("IdentityService.EnrollDevice() error = unexpected failed enrollment")
				}
			}
		})
	}
}

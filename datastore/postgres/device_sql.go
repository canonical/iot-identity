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

package postgres

const createDeviceTableSQL string = `
	CREATE TABLE IF NOT EXISTS device (
		id                serial primary key not null,
		device_id         varchar(200) not null unique,
		org_id            varchar(200) not null,
		brand             varchar(200) not null,
		model             varchar(200) not null,
		serial_number     varchar(200) not null,
		cred_key          text not null,
		cred_cert         text not null,
		cred_mqtt         varchar(200) not null,
		cred_port         varchar(200) not null,

		store_id          varchar(200) default '',
		device_key        text default '',
		status            int default 1,
        device_data       text default '',

        UNIQUE (device_id),
        UNIQUE (brand, model, serial_number)
	)
`

const createDeviceIDIndexSQL = "CREATE INDEX IF NOT EXISTS device_id_idx ON device (device_id)"

const createDeviceBMSIndexSQL = "CREATE INDEX IF NOT EXISTS bms_idx ON device (brand, model, serial_number)"

const createDeviceSQL = `
insert into device (device_id, org_id, brand, model, serial_number, cred_key, cred_cert, cred_mqtt, cred_port, device_data)
values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`

const getDeviceSQL = `
select device_id, org_id, brand, model, serial_number, cred_key, cred_cert, cred_mqtt, cred_port, store_id, device_key, status, device_data
from device
where brand=$1 and model=$2 and serial_number=$3`

const getDeviceByIDSQL = `
select device_id, org_id, brand, model, serial_number, cred_key, cred_cert, cred_mqtt, cred_port, store_id, device_key, status, device_data
from device
where device_id=$1`

const enrollDeviceSQL = `
update device
set store_id=$4, device_key=$5, status=$6
where brand=$1 and model=$2 and serial_number=$3
`

const updateDeviceSQL = `
update device
set status=$2, device_data=$3
where device_id=$1
`

const listDeviceSQL = `
select device_id, org_id, brand, model, serial_number, cred_cert, cred_mqtt, cred_port, store_id, device_key, status, device_data
from device
where org_id=$1`

// Add the device_data field to store a base64-encoded file
const alterDeviceAddDeviceData = "ALTER TABLE device ADD COLUMN device_data TEXT DEFAULT ''"

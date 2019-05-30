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

const createOrganizationTableSQL string = `
	CREATE TABLE IF NOT EXISTS organization (
		id               varchar(200) primary key not null,
		org_id           varchar(200) not null unique,
		name             varchar(200) not null,
		country_name     varchar(200) default '',
		root_cert         text not null,
		root_key          text not null,
        UNIQUE (org_id)
	)
`

const createOrganizationSQL = `
insert into organization (org_id, name, country_name, root_cert, root_key)
values ($1,$2,$3,$4,$5) RETURNING id`

const getOrganizationSQL = `
select id, org_id, name, country_name, root_cert, root_key
from organization
where org_id=$1`

const getOrganizationByNameSQL = `
select id, org_id, name, country_name, root_cert, root_key
from organization
where name=$1`

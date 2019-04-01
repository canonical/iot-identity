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
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Router returns the application router
func (wb IdentityService) Router() *mux.Router {
	// Start the web service router
	router := mux.NewRouter()

	router.Handle("/v1/organization", Middleware(http.HandlerFunc(wb.RegisterOrganization))).Methods("POST")
	router.Handle("/v1/device", Middleware(http.HandlerFunc(wb.RegisterDevice))).Methods("POST")
	router.Handle("/v1/device/enroll", Middleware(http.HandlerFunc(wb.EnrollDevice))).Methods("POST")

	return router
}

// Logger Handle logging for the web service
func Logger(start time.Time, r *http.Request) {
	log.Printf(
		"%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		time.Since(start),
	)
}

// Middleware to pre-process web service requests
func Middleware(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the request
		Logger(start, r)

		inner.ServeHTTP(w, r)
	})
}

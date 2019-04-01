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
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"log"
	"math/big"
	"time"

	"github.com/CanonicalLtd/iot-identity/domain"
)

// CreateClientCert creates a signed client certificate
func CreateClientCert(org *domain.Organization) ([]byte, []byte, error) {
	// Load the keypair from the PEM encoded data
	keyPair, err := tls.X509KeyPair(org.RootCert, org.RootKey)
	if err != nil {
		log.Println("Error parsing root keypair:", err)
		return nil, nil, err
	}

	// Parse the root certificate
	ca, err := parseRootCertificate(org.RootCert)
	if err != nil {
		log.Println("Error parsing root certificate:", err)
		return nil, nil, err
	}

	// Create a root certificate template
	serial, err := randomNumber()
	if err != nil {
		return nil, nil, err
	}
	cliTpl := clientTemplate(org, serial)

	// Generate a private key
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	pub := &privateKey.PublicKey

	// Sign the certificate
	cert, err := x509.CreateCertificate(rand.Reader, cliTpl, ca, pub, keyPair.PrivateKey)
	if err != nil {
		log.Println("Error creating client certificate:", err)
		return nil, nil, err
	}

	// Create plain text PEM for certificate
	certPEM := certToPEM(cert)

	// Create plain text PEM for key
	keyPEM := keyToPEM(privateKey)

	return keyPEM, certPEM, err
}

func clientTemplate(org *domain.Organization, serial *big.Int) *x509.Certificate {
	// Prepare certificate
	return &x509.Certificate{
		SerialNumber: serial,
		Subject: pkix.Name{
			Organization: []string{org.Name},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
}

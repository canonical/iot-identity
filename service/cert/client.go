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
	"time"

	"github.com/bugraaydogar/iot-identity/domain"
)

// CreateClientCert creates a signed client certificate
func CreateClientCert(org *domain.Organization, certsPath, deviceID string) ([]byte, []byte, error) {
	// Get the parsed CA from the filesystem
	caKeyPair, caTemplate, err := getCertificateAuthority(certsPath)
	if err != nil {
		return nil, nil, err
	}

	template := clientTemplate(org.Name, deviceID)
	privateKey, cert, err := createCertificate(template, caTemplate, caKeyPair)

	// Create plain text PEM for certificate
	certPEM := certToPEM(cert)

	// Create plain text PEM for key
	keyPEM := keyToPEM(privateKey)

	return keyPEM, certPEM, err
}

func createCertificate(template, parentTemplate *x509.Certificate, keyPair tls.Certificate) (*rsa.PrivateKey, []byte, error) {
	// Generate a private key
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	pub := &privateKey.PublicKey

	// Sign the certificate
	cert, err := x509.CreateCertificate(rand.Reader, template, parentTemplate, pub, keyPair.PrivateKey)
	if err != nil {
		log.Println("Error creating client certificate:", err)
		return nil, nil, err
	}
	return privateKey, cert, nil
}

func clientTemplate(name, deviceID string) *x509.Certificate {
	serial, err := randomNumber()
	if err != nil {
		return nil
	}

	// Prepare certificate
	return &x509.Certificate{
		SerialNumber: serial,
		Subject: pkix.Name{
			CommonName:   deviceID,
			Organization: []string{name},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
}

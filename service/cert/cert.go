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
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"time"
)

// CreateRootKeyCA generates the root key and certificate
func CreateRootKeyCA(org, country string) ([]byte, []byte, error) {
	// Create a root certificate template
	serial, err := randomNumber()
	if err != nil {
		return nil, nil, err
	}
	tpl := rootTemplate(org, country, serial)

	// Generate RSA private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	publicKey := &privateKey.PublicKey

	// Create a self-signed certificate. template = parent
	var parent = tpl
	ca, err := x509.CreateCertificate(rand.Reader, tpl, parent, publicKey, privateKey)

	// Create plain text PEM for CA
	caPEM := certToPEM(ca)

	// Create plain text PEM for key
	keyPEM := keyToPEM(privateKey)

	return keyPEM, caPEM, err
}

func rootTemplate(org, country string, serial *big.Int) *x509.Certificate {
	return &x509.Certificate{
		IsCA:                  true,
		BasicConstraintsValid: true,
		SubjectKeyId:          []byte{1, 2, 3},
		SerialNumber:          serial,
		Subject: pkix.Name{
			Organization: []string{org},
			Country:      []string{country},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(10, 0, 0),
		// see http://golang.org/pkg/crypto/x509/#KeyUsage
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}
}

func randomNumber() (*big.Int, error) {
	return rand.Int(rand.Reader, big.NewInt(8594))
}

func parseRootCertificate(rootCert []byte) (*x509.Certificate, error) {
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(rootCert)
	if !ok {
		return nil, fmt.Errorf("error using the root certificate")
	}

	block, _ := pem.Decode(rootCert)
	if block == nil {
		panic("failed to parse certificate PEM")
	}
	return x509.ParseCertificate(block.Bytes)
}

func certToPEM(c []byte) []byte {
	pemCA := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: c,
	}
	return pem.EncodeToMemory(pemCA)
}

func keyToPEM(key *rsa.PrivateKey) []byte {
	pemKey := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	return pem.EncodeToMemory(pemKey)
}

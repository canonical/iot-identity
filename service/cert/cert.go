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
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"path"
)

const (
	rootCA    = "ca.crt"
	rootCAKey = "ca.key"
)

// getCertificateAuthority loads the root certificate and key from the filesystem
func getCertificateAuthority(certsPath string) (tls.Certificate, *x509.Certificate, error) {
	ca := path.Join(certsPath, rootCA)
	key := path.Join(certsPath, rootCAKey)
	caKeyPair, err := tls.LoadX509KeyPair(ca, key)
	if err != nil {
		return tls.Certificate{}, nil, fmt.Errorf("cannot read root CA: %v", err)
	}
	caBytes, err := ioutil.ReadFile(ca)
	if err != nil {
		return tls.Certificate{}, nil, fmt.Errorf("cannot read root CA: %v", err)
	}
	caTemplate, err := parseRootCertificate(caBytes)
	if err != nil {
		return tls.Certificate{}, nil, fmt.Errorf("cannot read root CA: %v", err)
	}
	return caKeyPair, caTemplate, err
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
		return nil, fmt.Errorf("failed to parse certificate PEM")
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

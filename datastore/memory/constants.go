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

package memory

// CertPEM is a test CA
const CertPEM = `-----BEGIN CERTIFICATE-----
MIICrDCCAZQCCQCPLRGxNMKMKDANBgkqhkiG9w0BAQsFADAXMRUwEwYDVQQKDAxF
eGFtcGxlIEluYy4wIBcNMTkwMzI2MTYwMTA0WhgPMzAxODA3MjcxNjAxMDRaMBcx
FTATBgNVBAoMDEV4YW1wbGUgSW5jLjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC
AQoCggEBAMAuBU0Musn8o8hDI/xZZGDXr+LAKumx2Z7iIAJwTkSRF+c4yjV5FPd8
JDZTqyR9m1D12jvaG2cK78wJBeQdvsnJY9ARYYc7FuKBZbO3lm3pWaswMINCJdj5
XMVBaegrdKMlDLSXD2w0rE+Qh2kzEKYC3GHE4y0rQxaLJBIw0EgxO6pK2z/K4N6J
A9rqRPtfVfLvvAzVnZpPRraFhViNrqBIZRZiXOqTl2iBHPEmiBOXWx0ZdeLjgPO1
5/l0iNOUlgXkCX2Cn3sxx0aBY7Q57+VLB9ODwUnRgft9Usxpjf6QRSpvT4y90Tgq
tSucFCvS6GahcMs8mEDl6f8YVCg300UCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEA
P8FGaZLrU62/9MF0ObQbToE+e78gFR8Dx53a4WbIFWzzGx04EBuebqHBmRt4jjOP
B4UrQGfW1HUHm0Q3QtL7g2iO1TroGQIX+qVPd6+Aa4Kcq/zjFMK0p/Ikltp/KA6s
Sl0+CeyIgjlbiL05Rx+STdbMH4sWfaqYj242zy6hc3yXKH9wcUPmOXNvcCGIGAks
gJ0Q2lhKge4P8XdOta2aGaqSLT94UV8xlaCRlvh4jBxoRoH8cJjeHrHX2gv9X2mD
zuz5V4jcw1sQXfoCI9Lyd2VORRLThpWvCv5dHyTHMEqHKOjMcSu/cAhiHGtxi8BV
O/KPyd6tIUh+upwuVShqgw==
-----END CERTIFICATE-----`

// RootPEM is a test CA key
const RootPEM = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAwC4FTQy6yfyjyEMj/FlkYNev4sAq6bHZnuIgAnBORJEX5zjK
NXkU93wkNlOrJH2bUPXaO9obZwrvzAkF5B2+yclj0BFhhzsW4oFls7eWbelZqzAw
g0Il2PlcxUFp6Ct0oyUMtJcPbDSsT5CHaTMQpgLcYcTjLStDFoskEjDQSDE7qkrb
P8rg3okD2upE+19V8u+8DNWdmk9GtoWFWI2uoEhlFmJc6pOXaIEc8SaIE5dbHRl1
4uOA87Xn+XSI05SWBeQJfYKfezHHRoFjtDnv5UsH04PBSdGB+31SzGmN/pBFKm9P
jL3ROCq1K5wUK9LoZqFwyzyYQOXp/xhUKDfTRQIDAQABAoIBADwDoyAmo4ZEYRk+
7lP1zoT3ljOncz87jQwy7XAVhjufW+mXMH52a3fFysE0a7OfjgtAW4BpYjlRjwUW
pEJSj6wQOh1V8DD84O6nHg17fXyhbKErEVtMIumZJcFr5hjcyTXRciBLNEPERzMp
nT/a9I4DQrM9evw2EGNP3FnZ6JFCU/Of91PSI4jDFEr9y0oOs54PcN8kgZOuRvaB
vPXhwoXIVm1NZuaOSC8T793jcWGqr2pRaMiGtc5yPD/1dPE2YrYMQ0QbT95Ve39G
fI9bg7tIzCIyaT9GbuiOvTBrcCgxnhwrjkgiViGsEqRXXrpY89srLvCJS1BDou4A
MX6y7NkCgYEA6ky3PKWxmp0tT2RGt81PJ5e2+hJI0zHqrAzKVXj9M4ENmlh8zTU4
Rxym4Cll+y5LYlxEwoqEynd4sRahbhmxD30ILdYytckG2ZYIKwztj0D0Cf4+ZTP/
Wsu6YMEhTTMA2puJOQ1IyvuGm9k03nXwVLACsdg8TpeMKjPHRv+JSU8CgYEA0fqi
kV7daqTX86F3kWJ5Yc6WXC1EdQwOgyMa+/FK4NjlVLcxvGO2CptOJ6ScnCRqAPJB
HGSPRFFb+grBblmU0psvJgZYSz34q9W5ZGkOC9lYLLPkLNtcJSXViU7YeDAIHrWB
HPsvGSZSvUde1zl0OVmeWCHnv3KThhWUDbK4jSsCgYBtKjdJy423kzoUPo1wf/k8
YkS/uRszQ5Oqe/8d2dRnVd7Hpijn178T6vaZhNBeOtCm0IS8+5spVobmQ7wNN202
4TOZX66a4kINyQifPlPFJidOLKZXsuVsIXYCNJnWhUgFkuhZq6XZ5V1vacFnUR4b
5zIqOKzIlXWaCCv2GYOWowKBgQCUYfEPmWIOQn74g1njOxtbqolGihaeP+7hbKVc
9J5dVeh5fRuAbVXvGOCZ0xF4paLjGE46qjUzqeq9P2yBdnxcd51R2Zn7UcewZk+k
TTjH1scgj97mc/0hoyLK7RS7mfWi/dBHkpktxI8jgpPas5cWD+Z9kTgbafQmBImj
RHB2EQKBgQCEzEOLvic4dVUMPCXjw5dwjwytSqFAN2EVvVBVS83vhT26PgHiCd5N
qd6F62HNpV8FkRH8sUsiXAoVU+0hKp9mjilsFt6cP0jb+5/TGu8V2MHOrgmUmmVo
ZYUswIvPCkLb+jQzhG0JkQwEk0WWHYoXYINk+aYMYXQdvnXqJYxPLQ==
-----END RSA PRIVATE KEY-----`

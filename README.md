[![Build Status][travis-image]][travis-url]
[![Go Report Card][goreportcard-image]][goreportcard-url]
[![codecov][codecov-image]][codecov-url]
# IoT Identity Service

Managing the identity, ownership, credentials and authorization of an IoT device plays a crucial role in the security story. Those details need to be managed as the device goes through its lifecycle - from the manufacturer, distributor, system integrator, to end customer; from commissioning, repurposing to decommissioning the device.

The Identity service plays the role of managing these assets and enabling the connected systems to communicate with secure credentials.

The Identity Service is primarily in focus when the new device comes online. The device will be preconfigured to connect to the Identity Service, providing its Model and Serial assertions. The Identity Service registry will contain the primary ownership details for the device (customer name, store ID) and generates certificates and credentials for the device.

## Build
The project is using go modules.
Development has been done on minimum Go version 1.12.1.

```
$ cd iot-identity
$ go mod tidy
$ go build cmd/identity/main.go
```

## Run
```
go run cmd/identity/main.go
  -configdir string
        Directory path to the config file (default ".")
  -datasource string
        The data repository data source
  -driver string
        The data repository driver (default "memory")
  -mqttport string
        Port of the MQTT broker (default "8883")
  -mqtturl string
        URL of the MQTT broker (default "mqtt.example.com")
  -port string
        The port the service listens on (default "8030")
```

The service listens on 8030 by default.

## Contributing
Before contributing you should sign [Canonical's contributor agreement][1],
it’s the easiest way for you to give us permission to use your contributions.

[travis-image]: https://travis-ci.org/canonical/iot-identity.svg?branch=master
[travis-url]: https://travis-ci.org/canonical/iot-identity
[goreportcard-image]: https://goreportcard.com/badge/github.com/canonical/iot-identity
[goreportcard-url]: https://goreportcard.com/report/github.com/canonical/iot-identity
[codecov-url]: https://codecov.io/gh/canonical/iot-identity
[codecov-image]: https://codecov.io/gh/canonical/iot-identity/branch/master/graph/badge.svg
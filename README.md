# IoT Identity Service

Managing the identity, ownership, credentials and authorization of an IoT device plays a crucial role in the security story. Those details need to be managed as the device goes through its lifecycle - from the manufacturer, distributor, system integrator, to end customer; from commissioning, repurposing to decommissioning the device.

The Identity service plays the role of managing these assets and enabling the connected systems to communicate with secure credentials.

The Identity Service is primarily in focus when the new device comes online. The device will be preconfigured to connect to the Identity Service, providing its Model and Serial assertions. The Identity Service registry will contain the primary ownership details for the device (customer name, store ID) and generates certificates and credentials for the device.

## Build
The project uses vendorized dependencies using `govendor`.
Development has been done on minimum Go version 1.12.1.

```
$ go get github.com/CanonicalLtd/iot-identity
$ cd iot-identity
$ ./get-deps.sh
$ go build -mod=vendor ./...
```

## Run
```
go run -mod=vendor cmd/identity/main.go
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


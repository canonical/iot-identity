FROM golang:1.13 as builder1
COPY . ./src/github.com/canonical/iot-identity
WORKDIR /go/src/github.com/canonical/iot-identity
RUN CGO_ENABLED=1 GOOS=linux go build -a -o /go/bin/identity -ldflags='-extldflags "-static"' cmd/identity/main.go

# Copy the built applications to the docker image
FROM ubuntu:18.04
WORKDIR /srv
RUN apt-get update
RUN apt-get install -y ca-certificates
COPY --from=builder1 /go/bin/identity /srv/identity

# Note: /srv/certs is handled by the k8s file

# Set params from the environment variables
ARG DRIVER="postgres"
ARG DATASOURCE="dbname=identity sslmode=disable"
ARG PORT="8030"
ARG MQTTURL="localhost"
ARG MQTTPORT="8883"
ARG CERTSDIR="/srv/certs"
ARG CONFIGDIR="/srv/config"
ENV DRIVER="${DRIVER}"
ENV DATASOURCE="${DATASOURCE}"
ENV PORT="${PORT}"
ENV MQTTURL="${MQTTURL}"
ENV MQTTPORT="${MQTTPORT}"
ENV CERTSDIR="${CERTSDIR}"
ENV CONFIGDIR="${CONFIGDIR}"

EXPOSE 8030
ENTRYPOINT /srv/identity -port $PORT -driver $DRIVER -datasource "${DATASOURCE}" -mqtturl $MQTTURL -mqttport $MQTTPORT -certsdir $CERTSDIR -configdir $CONFIGDIR
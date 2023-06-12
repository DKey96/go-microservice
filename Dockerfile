FROM golang:1.19 AS build-env

# Secure against running as root
RUN mkdir /dk/

WORKDIR /dk-app/
ADD . /dk-app/

# Compile the binary, we don't want to run the cgo resolver
RUN CGO_ENABLED=0 go build -buildvcs=false -o /dk-app/app .

# final stage
FROM alpine:3.8

WORKDIR /
COPY --from=build-env /dk-app/certs/docker.localhost.* /
COPY --from=build-env /dk-app/app /

ENV DK_CERT_FILE=/docker.localhost.cert
ENV DK_KEY_FILE=/docker.localhost.key
ENV DK_SERVICE_ADDR=":8080"

EXPOSE 8080

CMD ["/app"]
#!/usr/bin/env bash

# CA
openssl genrsa -out ca.key 2048
openssl req -new -key ca.key -out ca.csr
openssl req -new -x509 -days 365 -key ca.key -out ca.crt

# Server
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr
openssl x509 -req -days 3650 -in server.csr -out server.crt -CA ca.crt -CAkey ca.key -CAcreateserial

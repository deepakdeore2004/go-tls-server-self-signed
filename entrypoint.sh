#!/bin/bash

cd /app

openssl genrsa -out server.key 2048
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 36500 -subj "/CN=example.com\/emailAddress=admin@example.com/C=US/ST=Ohio/L=Columbus/O=Widgets Inc/OU=Some Unit"

/app/server

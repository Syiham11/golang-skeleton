#!/bin/bash
set -o allexport; source .env; set +o allexport
echo "Updating API documentation..."
swag init
go build -oinjection.javamifi.com
./injection.javamifi.com

#!/usr/bin/env bash
./migrate.sh -path migrations -database "postgres://postgres:aladen@localhost/restapi_dev?sslmode=disable" up
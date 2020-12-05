#!/usr/bin/env bash
./migrate.sh create -ext sql -dir migrations $1
chmod -R o+rwx migrations
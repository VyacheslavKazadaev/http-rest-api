#!/usr/bin/env bash
if [ $1 = "pull" ]
then
docker pull migrate/migrate:latest
else
docker run --rm -v $(pwd)/migrations:/migrations --network host migrate/migrate $*
fi
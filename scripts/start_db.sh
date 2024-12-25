#! /usr/bin/env bash

podman run \
  --name=brynhildr-pg \
  --env-file=.env \
  --publish=5432:5432 \
  --volume=pg_data:/var/lib/postgresql/data \
  --volume=./db/init.sql:/docker-entrypoint-initdb.d/init.sql:ro \
  --replace \
  docker.io/library/postgres:17
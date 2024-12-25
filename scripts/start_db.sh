#! /usr/bin/env bash

podman run \
  --name=brynhildr-pg \
  --env-file=.env \
  --publish=5432:5432 \
  docker.io/library/postgres:17
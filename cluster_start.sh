#! /usr/bin/env bash

podman pod create \
    --publish 5432:5432 \
    --publish 8080:8080 \
    brynhildr
podman build -f Containerfile -t brynhildr-app
podman create \
    --pod brynhildr \
    --name postgres \
    --volume $(pwd)/db/init.sql:/docker-entrypoint-initdb.d/init.sql:ro \
    --volume pg_data:/var/lib/postgresql/data \
    --env-file .env \
    --replace \
    docker.io/library/postgres:17
podman create \
    --pod brynhildr \
    --name webserver \
    --env-file .env \
    --replace \
    --requires postgres \
    --restart always \
    localhost/brynhildr-app:latest

echo "starting pod"
podman pod start brynhildr
#!/bin/sh

docker compose down
docker compose build < Dockerfile
docker rmi $(docker images -qa -f 'dangling=true')
docker compose -f ./docker-compose.yml up
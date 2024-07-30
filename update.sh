#!/bin/bash

apt update -y
apt install -y git docker.io docker-compose-plugin

git pull
docekr compose down
docker compose build
docker compose up -d
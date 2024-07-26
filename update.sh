#!/bin/bash

apt update -y
apt install -y git docker.io docker-compose-plugin

git pull

docker-compose up -d
#!/bin/bash

sudo install -m 0755 -d /etc/apt/keyrings
sudo install -m 0755 -d /etc/apt/keyrings

DISTRO=$(lsb_release -is)

# Make $DISTRO lowercase
DISTRO=$(echo "$DISTRO" | tr '[:upper:]' '[:lower:]')

sudo curl -fsSL https://download.docker.com/linux/$DISTRO/gpg -o /etc/apt/keyrings/docker.asc

sudo chmod a+r /etc/apt/keyrings/docker.asc

echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/$DISTRO \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt update -y
sudo apt install -y git docker.io docker-compose-plugin ca-certificates curl
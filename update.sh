#!/bin/bash

docekr compose down
git pull
docker compose build
docker compose up -d
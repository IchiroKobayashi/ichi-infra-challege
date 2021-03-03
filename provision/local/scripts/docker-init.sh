#!/bin/bash

# Initialization of Environment variables
source .env

# Directory Definition
FIRST_OUTPUT='~/docker-compose'
LAST_OUTPUT='~/usr/local/bin'

# Download the latest docker compose
DOCKER_COMPOSE_LATEST_VERSION=$(curl https://api.github.com/repos/docker/compose/releases/latest | jq .name -r)
wget -q "https://github.com/docker/compose/releases/download/${DOCKER_COMPOSE_LATEST_VERSION}/docker-compose-$(uname -s)-$(uname -m)" -O $FIRST_OUTPUT
echo $(docker-compose --version)

# Make a directory & Move the installed docker-compose to the appropriate folder
sudo mkdir -p $LAST_OUTPUT
sudo mv ${FIRST_OUTPUT} ${LAST_OUTPUT}/docker-compose

# Change Owner & Mode for the Docker Compose
sudo chown root:root ${LAST_OUTPUT}/docker-compose
sudo chmod +x ${LAST_OUTPUT}/docker-compose

# Run containers using docker-compose
cd ${HOME_DIR}/provision/local
${LAST_OUTPUT}/docker-compose up -d --remove-orphans

sleep 10
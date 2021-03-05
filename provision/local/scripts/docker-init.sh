#!/bin/bash

# Initialization of Environment variables
# source .env #don't run this before cd
echo "[$(date)] Provisioning virtual machine..."

# Directory Definition
FIRST_OUTPUT='~/docker-compose'
LAST_OUTPUT='/usr/local/bin'

# Install jq command if necessary
jq --version >/dev/null 2>&1
if [ $? -ne 0 ]; then
    sudo yum -y install epel-release
    sudo yum -y install jq --enablerepo=epel
fi

# Install wget command if necessary
wget --version
if [ $? -ne 0 ]; then
    sudo yum -y install wget
fi

# Download the latest Docker
echo "Download the latest Docker"
DOCKER_LATEST_VERSION=$(curl -L https://api.github.com/repos/docker/compose/releases/latest | jq .name -r)
echo "Downloading docker v${DOCKER_LATEST_VERSION}..."
echo "This may take a few minutes..."
sudo yum install -y yum-utils device-mapper-persistent-data lvm2
wait
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
wait
sudo yum install -y docker-ce docker-ce-cli containerd.io
wait

# Download the latest docker compose
echo "Download the latest docker compose"
DOCKER_COMPOSE_LATEST_VERSION=$(curl -L https://api.github.com/repos/docker/compose/releases/latest | jq .name -r)
echo "Downloading docker-compose v${DOCKER_COMPOSE_LATEST_VERSION}..."
echo "This may take a few minutes..."
wget -q "https://github.com/docker/compose/releases/download/${DOCKER_COMPOSE_LATEST_VERSION}/docker-compose-$(uname -s)-$(uname -m)" -O $FIRST_OUTPUT
RESULT="$?"
[ "$RESULT" -ne "0" ] && echo "Failed to download docker-compose." 1>&2 && exit "$RESULT"
echo "Dowload completed　$(docker-compose --version)"

# Make a directory & Move the installed docker-compose to the appropriate folder
echo "Make a directory & Move the installed docker-compose to the appropriate folder"
sudo mkdir -p $LAST_OUTPUT
sudo mv ${FIRST_OUTPUT} ${LAST_OUTPUT}/docker-compose

# Change Owner & Mode for the Docker Compose
echo "Change Owner & Mode for the Docker Compose"
sudo chown root:root ${LAST_OUTPUT}/docker-compose
sudo chmod +x ${LAST_OUTPUT}/docker-compose

# Run containers using docker-compose
echo "[$(date)] Executing startup script..."
cd /home/cwd/src/provision/local
docker-compose up -d --remove-orphans

sleep 10
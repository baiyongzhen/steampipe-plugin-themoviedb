#!/bin/bash

set -e

# Install Docker
curl -fsSL get.docker.com -o get-docker.sh
sh get-docker.sh
sudo usermod -aG docker vagrant

# Install Docker Compose
sudo  curl -SL https://github.com/docker/compose/releases/download/v2.14.2/docker-compose-linux-x86_64 -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# Install Golang
# golang version number
GO_VERSION=1.20
sudo apt-get install -y curl
sudo curl -fsSL "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz" | sudo tar Cxz /usr/local

cat >> /home/vagrant/.profile <<EOF
GOPATH=\\$HOME/go
PATH=/usr/local/go/bin:\\$PATH
export GOPATH PATH
EOF
source /home/vagrant/.profile


# Install Make
sudo apt install make
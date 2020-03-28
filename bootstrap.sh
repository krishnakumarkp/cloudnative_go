#!/usr/bin/env bash

# create project folder
sudo mkdir "/files"
sudo apt-get update
sudo apt-get upgrade

sudo apt install -y apt-transport-https ca-certificates curl software-properties-common

sudo curl -fsSL https://get.docker.com -o get-docker.sh

while [ ! -f ./get-docker.sh ]; do sleep 1; done

sudo sh ./get-docker.sh

sudo usermod -aG docker vagrant
#!/usr/bin/env bash

# Update remote package metadata
apt-get update
apt-get -y upgrade
apt-get install git

sudo curl -O https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
tar -xvf go1.8.linux-amd64.tar.gz
mv go /usr/local

echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.profile
echo "export GOROOT=/usr/local/go" >> ~/.profile
echo "export GOPATH=/vagrant" >> ~/.profile
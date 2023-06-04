#!/bin/bash
SECONDS=0

cd $HOME/app

msg() {
	# Echo message in green color
	echo -e "\033[1;32m$1\033[0m"
}

msg "Pulling from GitHub"
git pull

msg "Building Docker image"
sudo docker build --tag app .

msg "Stopping Docker container"
sudo docker stop app
sudo docker rm app

msg "Starting Docker container"
sudo docker run \
-d \
--name app \
--expose 443 \
-p 443:443 \
-v /etc/letsencrypt:/etc/letsencrypt \
-e SERVER_ENV=PROD \
app

msg "Pruning stale Docker images"
sudo docker image prune -f
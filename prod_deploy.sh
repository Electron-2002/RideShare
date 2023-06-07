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

msg "Stopping containers"
sudo docker compose down

msg "Starting containers"
sudo docker compose up -d

msg "Pruning stale Docker images"
sudo docker image prune -f

duration=$SECONDS

echo
msg "Deployment completed successfully in $(($duration / 60)) seconds."
msg "Press Enter to exit."
read
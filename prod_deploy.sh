#!/bin/bash
SECONDS=0

cd $HOME/app

msg() {
	# Echo message in green color
	echo -e "\033[1;32m$1\033[0m"
}

msg "Stopping the application"
sudo pkill app

msg "Pulling the latest code from the git repository"
git pull

msg "Building the application"
go build

msg "Starting the application"
nohup sudo ./app &>/dev/null &

duration=$SECONDS

echo
msg "Deployment completed successfully in $(($duration / 60)) seconds."
msg "Press Enter to exit."
read

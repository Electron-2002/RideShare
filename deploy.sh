#!/bin/bash
sshcmd="ssh -t satej@rideshare.satejbidvai.tech"
$sshcmd screen -S "deployment" /home/satej/app/prod_deploy.sh

#!/bin/bash
sshcmd="ssh -t satej@rideshare.satejbidvai.tech"
$sshcmd screen -S "deployment" /home/satej/app/deploy_prod.sh

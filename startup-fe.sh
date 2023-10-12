#!/bin/bash
if [ ! -f /var/run/my_script_ran_before ]; then
    # Mark that the script has run before
    sudo touch /var/run/my_script_ran_before

    # Execute the desired script
    cd /
    apt-get -y update
    apt-get -y install pip
    apt-get -y install git
    apt-get -y install npm
    cd /
    git clone https://github.com/yilong100/GoApp.git
    cd /
    cd GoApp/frontend/react-app/
    npm install
    npm start &
fi
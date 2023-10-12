#!/bin/bash
if [ ! -f /var/run/my_script_ran_before ]; then
    # Mark that the script has run before
    sudo touch /var/run/my_script_ran_before

    # Execute the desired script
    cd /
    wget https://go.dev/dl/go1.21.3.linux-amd64.tar.gz
    tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz
    mkdir /root/go
    export GOPATH=/root/go
    export GOCACHE=/root/go/cache
    export PATH=${PATH}:/usr/local/go/bin:${GOPATH}/bin
    apt-get -y update
    apt-get -y install pip
    apt-get -y install git
    apt-get -y install golang
    apt-get -y install gnupg curl
    curl -fsSL https://pgp.mongodb.com/server-7.0.asc | gpg -o /usr/share/keyrings/mongodb-server-7.0.gpg --dearmor
    echo "deb [ signed-by=/usr/share/keyrings/mongodb-server-7.0.gpg ] http://repo.mongodb.org/apt/debian bullseye/mongodb-org/7.0 main" | tee /etc/apt/sources.list.d/mongodb-org-7.0.list
    apt-get -y update
    apt-get install -y mongodb-org
    echo "mongodb-org hold" | dpkg --set-selections
    echo "mongodb-org-database hold" | dpkg --set-selections
    echo "mongodb-org-server hold" | dpkg --set-selections
    echo "mongodb-mongosh hold" | dpkg --set-selections
    echo "mongodb-org-mongos hold" | dpkg --set-selections
    echo "mongodb-org-tools hold" | dpkg --set-selections
    systemctl start mongod
    systemctl daemon-reload
    systemctl enable mongod
    cd /
    git clone https://github.com/yilong100/GoApp.git
    cd /
    cd GoApp/frontend/react-app/
    npm install
    npm start &
fi
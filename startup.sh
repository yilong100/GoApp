#!/bin/bash
$HOME=/
cd
sudo apt-get update
y
sudo apt-get install pip
y
sudo apt-get install git
y
sudo apt-get install golang
y
sudo apt-get install gnupg curl
y
curl -fsSL https://pgp.mongodb.com/server-7.0.asc | sudo gpg -o /usr/share/keyrings/mongodb-server-7.0.gpg --dearmor
y
echo "deb [ signed-by=/usr/share/keyrings/mongodb-server-7.0.gpg ] http://repo.mongodb.org/apt/debian bullseye/mongodb-org/7.0 main" | sudo tee /etc/apt/sources.list.d/mongodb-org-7.0.list
y
sudo apt-get update
y
sudo apt-get install -y mongodb-org
y
echo "mongodb-org hold" | sudo dpkg --set-selections
echo "mongodb-org-database hold" | sudo dpkg --set-selections
echo "mongodb-org-server hold" | sudo dpkg --set-selections
echo "mongodb-mongosh hold" | sudo dpkg --set-selections
echo "mongodb-org-mongos hold" | sudo dpkg --set-selections
echo "mongodb-org-tools hold" | sudo dpkg --set-selections
sudo systemctl start mongod
sudo systemctl daemon-reload
sudo systemctl enable mongod
cd
git clone https://github.com/yilong100/GoApp.git
cd GoApp/backend
go build
./GoPractice &
cd
cd GoApp/frontend/react-app/
npm install
npm start &
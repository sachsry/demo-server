#!/bin/bash
echo "Hello World"
yum update -y
yum install golang -y
export ENV=dev
mkdir /home/code
cd /home/code
git clone https://github.com/sachsry/demo-server.git
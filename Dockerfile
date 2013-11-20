from ubuntu:12.04
run echo "deb http://archive.ubuntu.com/ubuntu precise main universe" > /etc/apt/sources.list
run apt-get update -y
run apt-get install -yq git golang
env GOPATH /app
run go get github.com/rochacon/ifconfigme
workdir /app
env PORT 8000
expose 8000
cmd /app/bin/ifconfigme

from rochacon/golang
run apt-get install -yq git && apt-get clean
env GOPATH /app
run go get github.com/rochacon/ifconfigme
workdir /app
env PORT 80
expose 80
cmd /app/bin/ifconfigme

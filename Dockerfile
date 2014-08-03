FROM rochacon/golang
RUN DEBIAN_FRONTEND=noninteractive apt-get install -yq git && apt-get clean
ENV GOPATH /app
RUN go get github.com/rochacon/ifconfigme
WORKDIR /app
ENV PORT 80
EXPOSE 80
CMD /app/bin/ifconfigme

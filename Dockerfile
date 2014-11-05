FROM golang:1.3
ENV GOPATH /app
ADD . /app/src/github.com/rochacon/ifconfigme
RUN go install -v github.com/rochacon/ifconfigme
WORKDIR /app
ENV PORT 80
EXPOSE 80
CMD /app/bin/ifconfigme

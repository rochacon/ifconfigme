FROM rochacon/golang
ENV GOPATH /app
ADD . /app/src/github.com/rochacon/ifconfigme
RUN go install -v github.com/rochacon/ifconfigme
WORKDIR /app
ENV PORT 80
EXPOSE 80
CMD /app/bin/ifconfigme

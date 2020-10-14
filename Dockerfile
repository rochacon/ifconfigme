FROM golang:1.15 AS build
WORKDIR /ifconfigme/
COPY . .
ENV CGO_ENABLED=0
RUN go build -v

FROM scratch
COPY --from=build /ifconfigme/ifconfigme /ifconfigme
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT ["/ifconfigme"]
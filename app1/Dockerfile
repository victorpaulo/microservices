# build stage
FROM golang:latest AS build-env
ADD . /src
RUN cd /src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app1

# final stage
FROM alpine
ENV APP2="http://192.168.0.4:8082/bar"
EXPOSE 8081
COPY --from=build-env /src/app1 .
ENTRYPOINT ./app1
# build stage
FROM golang:latest AS build-env
ADD . /src
RUN cd /src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app2

# final stage
FROM alpine
ENV APP1="http://192.168.0.4:8081/foo"
EXPOSE 8082
COPY --from=build-env /src/app2 .
ENTRYPOINT ./app2
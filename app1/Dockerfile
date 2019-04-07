FROM golang:latest
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
EXPOSE 8081
ENV APP2="http://localhost:8082/bar"
RUN go build -o app1 . 
CMD ["/app/app1"]
FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
EXPOSE 8082
ENV APP1="http://localhost:8081/foo"
RUN go build -o app2 . 
CMD ["/app/app2"]
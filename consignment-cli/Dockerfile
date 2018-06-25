FROM golang:1.10.3 as build
WORKDIR /go/src/github.com/tsauvajon/microservices/consignment-service
COPY . .
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:3.7
WORKDIR /app

ADD consignment.json .
ADD consignment-cli .
RUN chmod +x /app/consignment-cli

CMD ["/app/consignment-cli"]

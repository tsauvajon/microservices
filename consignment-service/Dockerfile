FROM golang:1.10.3 as build
WORKDIR /go/src/github.com/tsauvajon/microservices/consignment-service
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:3.7
WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY --from=build /go/src/github.com/tsauvajon/microservices/consignment-service/consignment-service .
RUN ls /app
RUN chmod +x consignment-service
CMD ["/app/consignment-service"]

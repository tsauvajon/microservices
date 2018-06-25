build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/tsauvajon/microservices/consignment-service proto/consignment/consignment.proto
	docker build -t consignment-service .

run:
	docker run -p 50051:50051 consignment-service \
        -e MICRO_SERVER_ADDRESS=:50051 \
        -e MICRO_REGISTRY=mdns consignment-service

dev:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/tsauvajon/microservices/consignment-service proto/consignment/consignment.proto
	MICRO_SERVER_ADDRESS=:50051 \
        MICRO_REGISTRY=mdns \
		go run main.go

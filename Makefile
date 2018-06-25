install:
	go get -u google.golang.org/grpc
	go get -u github.com/micro/protobuf/{proto,protoc-gen-go}

build:
	$(MAKE) -C consignment-service build
	$(MAKE) -C consignment-cli build

run:
	$(MAKE) -C consignment-service run &
	$(MAKE) -C consignment-cli run &

stop:
	docker stop $$(docker ps -aq)

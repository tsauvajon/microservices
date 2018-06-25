build:
	CGO_ENABLED=0 GOOS=linux go build
	docker build -t consignment-cli .

run:
	docker run -e MICRO_REGISTRY=mdns consignment-cli

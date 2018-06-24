package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/tsauvajon/microservices/consignment-service/proto/consignment"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	port := os.Getenv("GRPC_PORT")
	host := os.Getenv("GRPC_HOST")

	if port == "" {
		port = ":50051"
	}

	if host == "" {
		host = "localhost"
	}

	address := host + port

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewShippingServiceClient(conn)

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	response, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list the consignments: %v", err)
	}
	for _, c := range response.Consignments {
		log.Println(c)
	}
}

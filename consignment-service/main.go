package main

import (
	"log"
	"net"
	"os"

	pb "github.com/tsauvajon/microservices/consignment-service/proto/consignment"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Repository : repo
type Repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// DummyRepository : Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type DummyRepository struct {
	consignments []*pb.Consignment
}

// Create : add a consignment
func (repo *DummyRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

// GetAll : return all the consignments
func (repo *DummyRepository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition.
type service struct {
	repo Repository
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	// Save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments := s.repo.GetAll()
	return &pb.Response{Consignments: consignments}, nil
}

func main() {
	repo := &DummyRepository{}

	port := os.Getenv("GRPC_PORT")

	if port == "" {
		port = ":50051"
	}

	// Set-up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	pb.RegisterShippingServiceServer(s, &service{repo})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

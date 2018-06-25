package main

import (
	"fmt"

	micro "github.com/micro/go-micro"
	pb "github.com/tsauvajon/microservices/consignment-service/proto/consignment"
	"golang.org/x/net/context"
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
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	// Save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	res.Consignments = s.repo.GetAll()
	return nil
}

func main() {
	repo := &DummyRepository{}

	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

package grpc

import (
	"context"
	"log"
	"net"

	"github.com/pavelanni/go-grpc-course/internal/rocket"
	rkt "github.com/pavelanni/protos/rocket/v1"
	"google.golang.org/grpc"
)

// RocketService defines the interface
type RocketService interface {
	GetRocketByID(ctx context.Context, id string) (rocket.Rocket, error)
	InsertRocket(ctx context.Context, rkt rocket.Rocket) (rocket.Rocket, error)
	DeleteRocket(ctx context.Context, id string) error
}

// Handler handles incoming gRPC requests
type Handler struct {
	RocketService RocketService
}

// New creates a new Handler
func New(rktService RocketService) Handler {
	return Handler{
		RocketService: rktService,
	}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Print("could not listen to port 50051")
		return err
	}

	grpcServer := grpc.NewServer()
	rkt.RegisterRocketServiceServer(grpcServer, &h)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %s\n", err)
		return err
	}

	return nil

}

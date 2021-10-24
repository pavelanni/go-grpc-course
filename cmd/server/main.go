package main

import (
	"log"

	"github.com/pavelanni/go-grpc-course/internal/db"
	"github.com/pavelanni/go-grpc-course/internal/rocket"
)

// Run is responsible for initializing and starting our gRPC server
func Run() error {
	rocketStore, err := db.New()
	if err != nil {
		return err
	}
	err = rocketStore.Migrate()
	if err != nil {
		log.Println("Failed to run migrations")
		return err
	}

	_ = rocket.New(rocketStore)
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

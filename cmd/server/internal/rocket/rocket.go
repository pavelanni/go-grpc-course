//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/pavelanni/go-grpc-course/internal/rocket Store

package rocket

import "context"

// Rocket contains the definitions of out rocket
type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

// Store defines the interface that we expect our database implementation to follow
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Service is responsible for updating the rocket inventory
type Service struct {
	Store Store
}

// New returns a new instance of our rocket service
func New(store Store) Service {
	return Service{
		Store: store,
	}
}

// GetRocketByID retrieves a rocket from the store using its ID
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// InsertRocket adds a new rocket to the store
func (s Service) InstertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// DeleteRocket removes a rocket from the store
func (s Service) DeleteRocket(id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}

package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pavelanni/go-grpc-course/internal/rocket"
	"github.com/pkg/errors"
)

type Store struct {
	db *sqlx.DB
}

// ConnectLoop tries to connect to the DB under given DSN using a give driver
// in a loop until connection succeeds. timeout specifies the timeout for the
// loop.
// I got it from here: https://alex.dzyoba.com/blog/go-connect-loop/ Thanks, Alex Dyoba!
func ConnectLoop(driver, DSN string, timeout time.Duration) (*sqlx.DB, error) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeoutExceeded := time.After(timeout)
	for {
		fmt.Println("trying to connect at ", time.Now())
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("db connection failed after %s timeout", timeout)

		case <-ticker.C:
			db, err := sqlx.Connect("postgres", DSN)
			if err == nil {
				fmt.Println("connected to db")
				return db, nil
			}
			log.Println(errors.Wrapf(err, "failed to connect to db %s", DSN))
		}
	}
}

// New returns a new store or error
func New() (Store, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbTable := os.Getenv("DB_TABLE")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUsername,
		dbPassword,
		dbTable,
		dbSSLMode,
	)

	db, err := ConnectLoop("postgres", connectionString, 3*time.Minute)
	if err != nil {
		return Store{}, err
	}
	return Store{
		db: db,
	}, nil
}

func (s Store) GetRocketByID(id string) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

func (s Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

func (s Store) DeleteRocket(id string) error {
	return nil
}

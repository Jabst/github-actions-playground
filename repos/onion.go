package repos

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type OnionBreed string

var (
	ErrEmptyBasket error = errors.New("the basket is empty")
)

var (
	Brown OnionBreed = "Brown"
	Red   OnionBreed = "Red"
)

type Onion struct {
	Layers uint
	Breed  OnionBreed
}

type OnionStore struct {
	Onions []Onion
	db     *pgxpool.Pool
}

func NewOnionStore(pool *pgxpool.Pool) OnionStore {
	os := OnionStore{
		Onions: make([]Onion, 0),
		db:     pool,
	}

	os.InitDB(context.Background())

	return os
}

func (s *OnionStore) AddOnion(o Onion) int {
	s.Onions = append(s.Onions, o)

	return len(s.Onions)
}

func (s *OnionStore) ShiftOnions() (*Onion, error) {
	if len(s.Onions) == 0 {
		return nil, ErrEmptyBasket
	}

	removedOnion := s.Onions[0]

	s.Onions = s.Onions[1:]
	return &removedOnion, nil
}

func (s *OnionStore) PopOnion() (*Onion, error) {
	if len(s.Onions) == 0 {
		return nil, ErrEmptyBasket
	}

	removedOnion := s.Onions[len(s.Onions)-1]

	s.Onions = s.Onions[:len(s.Onions)-1]
	return &removedOnion, nil
}

func (s OnionStore) InitDB(ctx context.Context) error {
	_, err := s.db.Exec(ctx, `
	DROP TABLE IF EXISTS onions;

	CREATE TABLE onions (
		id SERIAL PRIMARY KEY,
		layers INTEGER NOT NULL,
		breed TEXT
	);`)

	if err != nil {
		return fmt.Errorf("failed to init db, %w", err)
	}

	return nil
}

func (s OnionStore) InsertOnions(ctx context.Context) error {

	for _, elem := range s.Onions {
		_, err := s.db.Exec(ctx, `
			INSERT INTO onions(layers, breed)
			VALUES($1, $2)
		`, elem.Layers, elem.Breed)

		// simulates very heavy insert for testing purposes
		time.Sleep(time.Second * 5)

		if err != nil {
			return fmt.Errorf("failed to exec sql [insert] statement on %+v onion, %w", elem, err)
		}
	}

	return nil
}

func (s *OnionStore) GetOnions(ctx context.Context) error {
	rows, err := s.db.Query(ctx, `SELECT layers, breed FROM onions`)
	if err != nil {
		return fmt.Errorf("failed to query onions")
	}

	for rows.Next() {
		var (
			layers uint
			breed  string
		)

		err := rows.Scan(&layers, &breed)
		if err != nil {
			return fmt.Errorf("failed to scan, %w", err)
		}

		// simulates very heavy select for testing purposes
		time.Sleep(time.Second * 2)

		s.Onions = append(s.Onions, Onion{
			Layers: layers,
			Breed:  OnionBreed(breed),
		})
	}

	return nil
}

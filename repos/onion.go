package repos

import "errors"

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
}

func NewOnionStore() OnionStore {
	return OnionStore{
		Onions: make([]Onion, 0),
	}
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

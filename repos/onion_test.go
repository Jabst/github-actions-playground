package repos_test

import (
	"Jabst/github-actions-playground/repos"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func initOnion() repos.OnionStore {
	return repos.NewOnionStore()
}

func setupOnion() []repos.Onion {
	return []repos.Onion{
		{
			Layers: 1,
			Breed:  repos.Brown,
		},
		{
			Layers: 4,
			Breed:  repos.Brown,
		},
		{
			Layers: 3,
			Breed:  repos.Brown,
		},
		{
			Layers: 2,
			Breed:  repos.Brown,
		},
		{
			Layers: 5,
			Breed:  repos.Red,
		},
	}
}

func Test_AddOnion(t *testing.T) {
	onionStore := initOnion()

	onion := repos.Onion{
		Layers: uint(2),
		Breed:  repos.Brown,
	}

	result := onionStore.AddOnion(onion)

	if diff := cmp.Diff(result, 1); diff != "" {
		t.Fatalf("expected not same with actual result: %s", diff)
	}
}

func Test_AddOnion_2(t *testing.T) {
	onionStore := initOnion()

	onion := repos.Onion{
		Layers: uint(2),
		Breed:  repos.Brown,
	}

	onionStore.AddOnion(onion)
	result := onionStore.AddOnion(onion)

	if diff := cmp.Diff(result, 2); diff != "" {
		t.Fatalf("expected not same with actual result: %s", diff)
	}
}

func Test_ShiftOnion(t *testing.T) {
	onionStore := initOnion()

	onionsToAdd := setupOnion()

	onionStore.Onions = append(onionStore.Onions, onionsToAdd...)

	result, err := onionStore.ShiftOnions()

	if err != nil {
		t.Fatalf("error was thrown, %s", err)
	}

	if diff := cmp.Diff(result, &repos.Onion{Layers: 1, Breed: repos.Brown}); diff != "" {
		t.Fatalf("expected not same with actual result: %s", diff)
	}
}

func Test_ShiftOnion_ToFail(t *testing.T) {
	onionStore := initOnion()

	onionsToAdd := setupOnion()

	onionStore.Onions = append(onionStore.Onions, onionsToAdd...)

	onionStore.ShiftOnions()
	onionStore.ShiftOnions()
	onionStore.ShiftOnions()
	onionStore.ShiftOnions()
	onionStore.ShiftOnions()
	onionStore.ShiftOnions()
	onionStore.ShiftOnions()
	_, err := onionStore.ShiftOnions()

	if !errors.Is(err, repos.ErrEmptyBasket) {
		t.Fatalf("error should be thrown as %s but is %s", repos.ErrEmptyBasket, err)
	}

}

func Test_PopOnion(t *testing.T) {
	onionStore := initOnion()

	onionsToAdd := setupOnion()

	onionStore.Onions = append(onionStore.Onions, onionsToAdd...)

	result, err := onionStore.PopOnion()

	if err != nil {
		t.Fatalf("error was thrown, %s", err)
	}

	if diff := cmp.Diff(result, &repos.Onion{Layers: 5, Breed: repos.Red}); diff != "" {
		t.Fatalf("expected not same with actual result: %s", diff)
	}
}

func Test_PopOnion_ToFail(t *testing.T) {
	onionStore := initOnion()

	onionsToAdd := setupOnion()

	onionStore.Onions = append(onionStore.Onions, onionsToAdd...)

	onionStore.PopOnion()
	onionStore.PopOnion()
	onionStore.PopOnion()
	onionStore.PopOnion()
	onionStore.PopOnion()
	onionStore.PopOnion()
	onionStore.PopOnion()
	_, err := onionStore.PopOnion()

	if !errors.Is(err, repos.ErrEmptyBasket) {
		t.Fatalf("error should be thrown as %s but is %s", repos.ErrEmptyBasket, err)
	}
}

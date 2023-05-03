package repos_test

import (
	"Jabst/github-actions-playground/repos"
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

func initOnionStore() repos.OnionStore {

	connString := "postgres://user:password@postgres:53/onions"

	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		panic(err)
	}

	store := repos.NewOnionStore(pool)

	_, err = pool.Exec(context.Background(), `delete from onions;
		ALTER SEQUENCE onions_id_seq RESTART WITH 1;
		INSERT INTO onions(layers, breed)
		VALUES (100, 'red'),
			(50, 'brown');`)
	if err != nil {
		panic(err)
	}

	return store
}

func Test_Integration_GetOnions(t *testing.T) {

	store := initOnionStore()

	err := store.GetOnions(context.Background())

	if err != nil {
		t.Fatalf("errored %s", err)
	}

	fmt.Printf("%+v\n", store.Onions)
}

func Test_Integration_InsertOnions(t *testing.T) {

	store := initOnionStore()

	store.AddOnion(repos.Onion{
		Layers: 1,
		Breed:  repos.Brown,
	})

	err := store.InsertOnions(context.Background())

	if err != nil {
		t.Fatalf("errored %s", err)
	}

	fmt.Printf("%+v\n", store.Onions)
}

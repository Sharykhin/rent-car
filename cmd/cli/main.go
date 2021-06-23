package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"Sharykhin/rent-car/di"
)

// TODO: This is used for simple tests. Remove it at the end
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = di.Init()
	if err != nil {
		log.Fatalf("failed to initialize di container: %v", err)
	}
	postgres := di.Container.PostgresConn
	err = postgres.Connect()
	defer postgres.Close()

	ctx := context.Background()
	decor(ctx, func(txCtx context.Context) error {
		a(txCtx, 10)

		return errors.New("ha ha ha")
	})

}

func a(ctx context.Context, in int) bool {
	f := ctx.Value("foo")
	fmt.Println("foo", f, in)

	return true
}

func b(s string) string {
	return s + "10"
}

func decor(ctx context.Context, fn func(context.Context) error) {
	fmt.Println("decor")
	newCtx := context.WithValue(ctx, "foo", "bar")
	err := fn(newCtx)
	if err != nil {
		fmt.Println("rollback")
	} else {
		fmt.Println("commit?")
	}

}

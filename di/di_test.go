package di

import (
	"log"
	"os"
	"testing"
)

func init() {
	err := os.Setenv("POSTGRES_URL", "test")
	if err != nil {
		log.Fatalf("failed to set tets env vatiable: %v", err)
	}
}

func TestInit(t *testing.T) {
	err := Init()
	if err != nil {
		t.Fatalf("expected err nil but got %v", err)
	}

	err = Init()

	if err != AlreadyInitializedError {
		t.Fatalf("expected err %v but got %v", AlreadyInitializedError, err)
	}
}

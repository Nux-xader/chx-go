package chx_test

import (
	"errors"
	"testing"

	"github.com/Nux-xader/chx-go"
)

const (
	serverAddr = "127.0.0.1:3800"
)

// TestNewClient tests the NewClient function.
func TestNewClient(t *testing.T) {
	t.Run("ValidAddress", func(t *testing.T) {
		client, err := chx.NewClient(serverAddr)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if client == nil {
			t.Fatal("expected client to be non-nil")
		}
		defer client.Close()
	})

	t.Run("InvalidAddress", func(t *testing.T) {
		_, err := chx.NewClient("127.0.0.1:9999")
		if err == nil {
			t.Fatal("expected an error, got nil")
		}
	})

	t.Run("DefaultAddress", func(t *testing.T) {
		client, err := chx.NewClient("") // "" defaults to serverAddr
		if err != nil {
			t.Fatalf("expected no error for default address, got %v", err)
		}
		defer client.Close()
	})
}

// TestGet tests the Get method.
func TestGet(t *testing.T) {
	client, err := chx.NewClient(serverAddr)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	t.Run("NotFound", func(t *testing.T) {
		_, err := client.Get("nonexistent")
		if !errors.Is(err, chx.ErrNotFound) {
			t.Errorf("expected ErrNotFound, got %v", err)
		}
	})

	t.Run("EmptyKey", func(t *testing.T) {
		_, err := client.Get("")
		if !errors.Is(err, chx.ErrNotFound) {
			t.Errorf("expected ErrNotFound for empty key, got %v", err)
		}
	})
}

// TestDelete tests the Delete method.
func TestDelete(t *testing.T) {
	client, err := chx.NewClient(serverAddr)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	t.Run("NotFound", func(t *testing.T) {
		err := client.Delete("nonexistent")
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("EmptyKey", func(t *testing.T) {
		err := client.Delete("")
		if !errors.Is(err, chx.ErrNotFound) {
			t.Errorf("expected ErrNotFound for empty key, got %v", err)
		}
	})
}

// TestConnectionClose tests behavior when the connection is closed.
func TestConnectionClose(t *testing.T) {
	client, err := chx.NewClient(serverAddr)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	// Close the client connection
	client.Close()

	_, err = client.Get("anykey")
	if err == nil {
		t.Error("expected an error when getting from a closed connection, got nil")
	}
}
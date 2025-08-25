# chx-go

`chx-go` is a lightweight, high-performance, and thread-safe Go client library for the `chx` key-value store server. It provides a simple and intuitive API for interacting with `chx`, making it easy to integrate into your Go applications.

## Features

*   **Lightweight and Fast:** The client is designed for minimal overhead and high performance.
*   **Thread-Safe:** The client can be safely used across multiple goroutines.
*   **Simple API:** The API is clean, intuitive, and easy to use.

## Installation

To install the `chx-go` library, use `go get`:

```bash
go get github.com/Nux-xader/chx-go
```

## Usage

The following example demonstrates how to start a `chx` server (assuming the binary is available), connect to it using `chx-go`, and perform basic `Set`, `Get`, and `Delete` operations.

```go
package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/Nux-xader/chx-go"
)

func main() {
	// Path to the chx server binary
	// IMPORTANT: Replace this with the actual path to your chx server executable.
	// You can download it from https://github.com/Nux-xader/chx/releases
	chxPath := "/media/nux/Dataxx/project/personal/chx/target/release/chx"

	// Start the chx server as a background process
	cmd := exec.Command(chxPath)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("Failed to start chx server: %v", err)
	}
	defer func() {
		if err := cmd.Process.Kill(); err != nil {
			log.Printf("Failed to kill server process: %v", err)
		}
	}()

	// Give the server a moment to start
	time.Sleep(1 * time.Second)

	// Connect to the server
	client, err := chx.NewClient(chx.DefaultAddr)
	if err != nil {
		log.Fatalf("Failed to connect to chx server: %v", err)
	}
	defer client.Close()

	// Perform Set operation
	err = client.Set("hello", "world")
	if err != nil {
		log.Fatalf("Set operation failed: %v", err)
	}
	fmt.Println("Set 'hello' to 'world'")

	// Perform Get operation
	value, err := client.Get("hello")
	if err != nil {
		log.Fatalf("Get operation failed: %v", err)
	}
	fmt.Printf("Get 'hello': %s\n", value)

	// Perform Delete operation
	err = client.Delete("hello")
	if err != nil {
		log.Fatalf("Delete operation failed: %v", err)
	}
	fmt.Println("Deleted 'hello'")

	// Verify deletion
	_, err = client.Get("hello")
	if err != nil {
		fmt.Printf("Get 'hello' after deletion: %v\n", err)
	}
}
```

## Contributing

Contributions are welcome! If you find a bug or have a feature request, please open an issue on GitHub.
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

### Importing the Library

To use the library, import the `chx-go` package in your Go code:

```go
import "github.com/Nux-xader/chx-go"
```

### Creating a New Client

To create a new client, use the `NewClient` function. If no address is provided, it will connect to the default address (`127.0.0.1:3800`).

```go
client, err := chx.NewClient("127.0.0.1:3800")
if err != nil {
    log.Fatalf("Failed to connect to chx server: %v", err)
}
defer client.Close()
```

### Setting a Value

To set a key-value pair, use the `Set` method.

```go
err := client.Set("mykey", "myvalue")
if err != nil {
    log.Printf("Failed to set key: %v", err)
}
```

### Getting a Value

To retrieve the value for a key, use the `Get` method.

```go
value, err := client.Get("mykey")
if err != nil {
    if errors.Is(err, chx.ErrNotFound) {
        fmt.Println("Key not found")
    } else {
        log.Printf("Failed to get key: %v", err)
    }
} else {
    fmt.Println("Found value:", value)
}
```

### Deleting a Value

To delete a key, use the `Delete` method.

```go
err := client.Delete("mykey")
if err != nil {
    log.Printf("Failed to delete key: %v", err)
}
```

### Error Handling

The library provides custom error types for more granular error handling.

*   `ErrNotFound`: Returned when a key is not found.
*   `ErrServer`: Returned for errors that occur on the `chx` server.

```go
value, err := client.Get("nonexistentkey")
if err != nil {
    var serverErr *chx.ErrServer
    if errors.As(err, &serverErr) {
        fmt.Println("Server error:", serverErr)
    } else if errors.Is(err, chx.ErrNotFound) {
        fmt.Println("Key not found")
    } else {
        fmt.Println("An unexpected error occurred:", err)
    }
}
```

### Closing the Connection

To close the connection to the `chx` server, use the `Close` method. It is recommended to use `defer` to ensure the connection is closed.

```go
defer client.Close()
```

## API Reference

### `type Client`

The `Client` struct is the primary entry point for interacting with the `chx` server.

### `func NewClient(address string) (*Client, error)`

`NewClient` creates a new `chx` client and connects to the server at the specified address. If the address is an empty string, it defaults to `"127.0.0.1:3800"`.

### `func (c *Client) Get(key string) (string, error)`

`Get` retrieves the value for a given key. It returns `ErrNotFound` if the key does not exist.

### `func (c *Client) Set(key, value string) error`

`Set` stores a key-value pair.

### `func (c *Client) Delete(key string) error`

`Delete` removes a key from the store.

### `func (c *Client) Close() error`

`Close` closes the connection to the `chx` server.

## Contributing

Contributions are welcome! If you find a bug or have a feature request, please open an issue on GitHub.
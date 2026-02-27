package chx

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"
)

// DefaultAddr is the default address for the chx server.
const DefaultAddr = "127.0.0.1:3800"

// Client is a client for the chx server.
// It is safe for concurrent use by multiple goroutines.
type Client struct {
	conn net.Conn
	mu   sync.Mutex
	addr string
}

// NewClient creates a new chx client.
// It connects to the chx server at the specified address.
// If address is empty, it uses DefaultAddr.
func NewClient(address string) (*Client, error) {
	if address == "" {
		address = DefaultAddr
	}
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return &Client{
		conn: conn,
		addr: address,
	}, nil
}

// Get retrieves the value for a key.
func (c *Client) Get(key string) (string, error) {
	if key == "" {
		return "", ErrNotFound
	}

	response, err := c.execute(fmt.Sprintf("G %s\n", key))
	if err != nil {
		return "", err
	}
	if response == "!" {
		return "", ErrNotFound
	}
	if strings.HasPrefix(response, "!e") {
		return "", &ErrServer{Err: errors.New(response[2:])}
	}
	if strings.HasPrefix(response, ">") {
		return response[1:], nil
	}
	return "", fmt.Errorf("invalid response: %s", response)
}

// Set sets the value for a key.
func (c *Client) Set(key, value string) error {
	response, err := c.execute(fmt.Sprintf("S %s %s\n", key, value))
	if err != nil {
		return err
	}
	if response == "!" {
		return nil
	}
	if strings.HasPrefix(response, "!e") {
		return &ErrServer{Err: errors.New(response[2:])}
	}
	return fmt.Errorf("invalid response: %s", response)
}

// Delete deletes a key.
func (c *Client) Delete(key string) error {
	if key == "" {
		return ErrNotFound
	}

	response, err := c.execute(fmt.Sprintf("D %s\n", key))
	if err != nil {
		return err
	}
	if response == "!" {
		return nil
	}
	if strings.HasPrefix(response, "!e") {
		return &ErrServer{Err: errors.New(response[2:])}
	}
	return fmt.Errorf("invalid response: %s", response)
}

func (c *Client) reconnect() error {
	if c.conn != nil {
		c.conn.Close()
	}
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *Client) execute(query string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// First attempt
	_, err := fmt.Fprint(c.conn, query)
	if err == nil {
		var response string
		response, err = c.readResponse()
		if err == nil {
			return response, nil
		}
	}

	// If failed (write or read), try reconnect and retry once
	if err := c.reconnect(); err != nil {
		return "", fmt.Errorf("reconnect failed: %w", err)
	}

	// Second attempt after reconnect
	_, err = fmt.Fprint(c.conn, query)
	if err != nil {
		return "", err
	}
	return c.readResponse()
}

// Close closes the connection to the chx server.
func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) readResponse() (string, error) {
	// Gunakan buffer besar (64KB) seperti implementasi Rust
	buffer := make([]byte, 65536)
	n, err := c.conn.Read(buffer)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(buffer[:n])), nil
}

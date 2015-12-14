package irc

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

var (
	// ErrInvalidServerURL is thrown when somebody attempts to create an irc object without a server.
	ErrInvalidServerURL = errors.New("server must be set")

	// ErrNilConnection is thrown when somebody writes to an unitilized connection
	ErrNilConnection = errors.New("You must Connect() before writing to an irc connection")
)

// IRC is a object holding irc connection information
type IRC struct {
	UseTLS     bool
	Server     string
	Output     io.Writer
	connection net.Conn
	Connected  bool
	timeout    time.Duration
}

// New creates a new irc connection
func New(server string, useTLS bool, output io.Writer, timeout time.Duration) (*IRC, error) {
	if server == "" {
		return nil, ErrInvalidServerURL
	}

	return &IRC{
		UseTLS:     useTLS,
		Server:     server,
		Output:     output,
		Connected:  false,
		connection: nil,
		timeout:    timeout,
	}, nil
}

// Write writes a message to the irc server.
func (i *IRC) Write(msg string) error {
	if i.connection == nil {
		return ErrNilConnection
	}
	return writeCon(i.connection, []byte(msg+"\r\n"), i.timeout)
}

// Connect blocks until the irc connection succeeds or times out
func (i *IRC) Connect() error {
	var err error
	i.connection, err = dial(i.Server, i.UseTLS, i.timeout)
	if err != nil {
		fmt.Fprintln(i.Output, err)
		return err
	}

	p := parser{i: i, outLock: &sync.Mutex{}}

	readLoop(i.connection, &p)
	i.Connected = true

	return nil
}

// closeConnection closes connection to server
func (i *IRC) closeConnection() error {
	return i.connection.Close()
}

// disconnect disconnects from server and cleans up
func (i *IRC) disconnect() {
	// Make sure connection is closed, we don't care if this errors
	i.closeConnection()
	i.connection = nil
}

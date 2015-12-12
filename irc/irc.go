package irc

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

// ErrInvalidConnection is thrown when somebody attempts to create an invalid connection.
var ErrInvalidConnection = errors.New("nick, user, and server must be set")

// IRC is a object holding irc connection information
type IRC struct {
	Nick       string
	User       string
	UseTLS     bool
	Server     string
	ServerPass string
	Output     io.Writer
	connection net.Conn
	connected  bool
	timeout    time.Duration
	readChan   chan bool
}

// New creates a new irc connection
func New(nick, user, server, serverPass string, useTLS bool, output io.Writer, timeout time.Duration) (*IRC, error) {
	if nick == "" || user == "" || server == "" {
		return nil, ErrInvalidConnection
	}

	return &IRC{
		Nick:       nick,
		User:       user,
		UseTLS:     useTLS,
		Server:     server,
		ServerPass: serverPass,
		Output:     output,
		connection: nil,
		connected:  false,
		timeout:    timeout,
	}, nil
}

// Write writes a message to the irc server.
func (i *IRC) Write(msg string) error {
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

	i.readChan = readLoop(i.connection, &p)

	if i.ServerPass != "" {
		if err = i.Write(passMsg(i.ServerPass)); err != nil {
			fmt.Fprintln(i.Output, err)
			return err
		}
	}

	if err = i.Write(nickMsg(i.Nick)); err != nil {
		fmt.Fprintln(i.Output, err)
		return err
	}

	if err = i.Write(userMsg(i.User, i.User)); err != nil {
		fmt.Fprintln(i.Output, err)
		return err
	}

	return nil
}

// Disconnect disconnects from server and cleans up
func (i *IRC) Disconnect() {
	i.Quit()
	time.Sleep(time.Second) // TODO: Wait for quit event
	i.readChan <- true
	<-i.readChan
}

package irc

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

// ErrWrite is thrown when not all data is written
var ErrWrite = errors.New("error writing all data")

func dial(address string, useTLS bool, timeout time.Duration) (net.Conn, error) {
	if useTLS {
		return tls.DialWithDialer(&net.Dialer{Timeout: timeout}, "tcp", address, nil)
	}

	return net.DialTimeout("tcp", address, timeout)
}

func writeCon(con net.Conn, msg []byte, timeout time.Duration) error {
	fmt.Println(string(msg))
	err := con.SetWriteDeadline(time.Now().Add(timeout))
	if err != nil {
		return err
	}

	n, err := con.Write(msg)

	if err != nil {
		return err
	}

	if n != len(msg) {
		return ErrWrite
	}

	return nil
}

func readLoop(con net.Conn, out io.Writer) chan bool {

	quit := make(chan bool)

	go func() {
		for {
			select {
			default:
				data := make([]byte, 512)
				_, err := con.Read(data)

				if err != nil {
					fmt.Print("error: ")
					fmt.Println(err)
				}

				fmt.Fprint(out, string(data))
			case <-quit:
				close(quit)
				return
			}
		}
	}()

	return quit
}

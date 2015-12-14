package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/chatterbox-irc/ircc/events"
	"github.com/chatterbox-irc/ircc/irc"
	"github.com/chatterbox-irc/ircc/parser"
)

const (
	connectionTimeout = time.Second * 10
)

var (
	server = flag.String("server", "", "Server address")
	tls    = flag.Bool("tls", false, "Connect using tls")
)

type connectionError struct {
}

func main() {
	flag.Parse()
	connect(os.Stdout, os.Stdin)
}

func connect(w io.Writer, r io.Reader) {
	ircc, err := irc.New(*server, *tls, w, 10*time.Second)
	if err != nil {
		fmt.Fprintln(w, events.ConnectionError(err.Error()))
		os.Exit(1)
	}

	err = ircc.Connect()

	if err != nil {
		os.Exit(2)
	}

	inReader := bufio.NewReader(os.Stdin)

	for {
		input, err := inReader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(w, events.InternalError(err.Error()))
		}

		status := parser.Parse(ircc, w, input)

		if status != -1 {
			os.Exit(status)
		}
	}
}

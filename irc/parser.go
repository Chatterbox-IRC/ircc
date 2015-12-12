package irc

import (
	"fmt"
	"strings"
	"sync"

	"github.com/chatterbox-irc/chatterbox/ircc/events"
)

const lineSplitter = "\r\n"

type parser struct {
	out     string
	outLock *sync.Mutex
	i       *IRC
}

// Write implements io.Writer to parse connection input
func (p *parser) Write(in []byte) (n int, err error) {
	p.outLock.Lock()
	p.out += string(in)
	lines := strings.Split(p.out, lineSplitter)
	if len(lines) > 1 {
		for _, line := range lines[:len(lines)-2] {
			p.i.ParseLine(line)
		}
	}
	p.out = lines[len(lines)-1]
	p.outLock.Unlock()
	return len(in), nil
}

// ParseLine parses output from IRC server.
func (i *IRC) ParseLine(line string) {
	msg := strings.Split(line, " ")

	if len(msg) < 2 {
		fmt.Fprintln(i.Output, events.InvalidMsgError())
		return
	}

	cmd := msg[1]

	switch cmd {
	case "001":
		parseWelcome(i, line)
	default:
		fmt.Fprintln(i.Output, line)
	}
}

func parseWelcome(i *IRC, line string) {
	// extract host from ':example.com 001 ...'
	server := strings.Split(strings.Split(line, " ")[0], ":")[1]

	split := strings.Split(line, " :")

	if len(split) < 2 {
		fmt.Fprintln(i.Output, events.InvalidMsgError())
		return
	}

	msg := split[1]

	fmt.Fprintln(i.Output, events.Connected(string(server), msg))
}

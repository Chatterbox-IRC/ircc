package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/chatterbox-irc/ircc/events"
	"github.com/chatterbox-irc/ircc/irc"
	"github.com/chatterbox-irc/pkg/validate"
)

// Parse parsses the input string and runs the command
// Returns -1 if successful, 0 to signal graceful exit, >0 to signify critical exit
func Parse(ircc *irc.IRC, w io.Writer, input string) int {
	cmd := make(map[string]interface{})

	if err := json.Unmarshal([]byte(input), &cmd); err != nil {
		fmt.Fprint(w, events.JSONError(err.Error()))
		return -1 // Not a critical error.
	}

	if _, ok := cmd["type"].(string); !ok {
		fmt.Fprint(w, events.JSONError("type field must exist and be a string"))
		return -1 // Not a critical error.
	}

	switch cmd["type"].(string) {
	case "quit":
		return quit(ircc, w)
	case "user":
		return user(ircc, w, input)
	case "nick":
		return nick(ircc, w, input)
	case "join":
		return join(ircc, w, input)
	case "part":
		return part(ircc, w, input)
	case "msg":
		return msg(ircc, w, input)
	default:
		fmt.Fprint(w, events.JSONError(fmt.Sprintf("unknown type %s", cmd["type"].(string))))
		return -1 // Not a critical error.
	}
}

func quit(ircc *irc.IRC, w io.Writer) int {
	ircc.Quit()
	start := time.Now()
	timeout := 2 * time.Second

	for {

		if ircc.Connected == false {
			return 0
		}

		if time.Since(start) > timeout {
			fmt.Fprintln(w, events.ConnectionError("unable to disconnect"))
			return 1
		}
	}
}

func user(ircc *irc.IRC, w io.Writer, input string) int {
	cmd := events.User{}

	if err := json.Unmarshal([]byte(input), &cmd); err != nil {
		fmt.Fprint(w, events.JSONError(err.Error()))
		return -1 // Not a critical error.
	}

	e := []validate.ValidationMsg{}
	e = append(e, validate.NotNil("user", cmd.User)...)
	e = append(e, validate.NotNil("name", cmd.Name)...)

	if len(e) > 0 {
		fmt.Fprint(w, events.ValidationError("user", e))
		return -1 // Not a critical error.
	}

	ircc.User(cmd.User, cmd.Name)
	return -1
}

func nick(ircc *irc.IRC, w io.Writer, input string) int {
	cmd := events.Nick{}

	if err := json.Unmarshal([]byte(input), &cmd); err != nil {
		fmt.Fprint(w, events.JSONError(err.Error()))
		return -1 // Not a critical error.
	}

	e := []validate.ValidationMsg{}
	e = append(e, validate.NotNil("nick", cmd.Nick)...)

	if len(e) > 0 {
		fmt.Fprint(w, events.ValidationError("nick", e))
		return -1 // Not a critical error.
	}

	ircc.Nick(cmd.Nick)
	return -1
}

func join(ircc *irc.IRC, w io.Writer, input string) int {
	cmd := events.Join{}

	if err := json.Unmarshal([]byte(input), &cmd); err != nil {
		fmt.Fprint(w, events.JSONError(err.Error()))
		return -1 // Not a critical error.
	}

	e := []validate.ValidationMsg{}
	e = append(e, validate.NotNil("channel", cmd.Channel)...)

	if len(e) > 0 {
		fmt.Fprint(w, events.ValidationError("join", e))
		return -1 // Not a critical error.
	}

	ircc.Join(cmd.Channel, cmd.Password)
	return -1
}

func part(ircc *irc.IRC, w io.Writer, input string) int {
	cmd := events.Part{}

	if err := json.Unmarshal([]byte(input), &cmd); err != nil {
		fmt.Fprint(w, events.JSONError(err.Error()))
		return -1 // Not a critical error.
	}

	e := []validate.ValidationMsg{}
	e = append(e, validate.NotNil("channel", cmd.Channel)...)

	if len(e) > 0 {
		fmt.Fprint(w, events.ValidationError("part", e))
		return -1 // Not a critical error.
	}

	ircc.Part(cmd.Channel)
	return -1
}

// TODO: Add CTCP support for newlines
// TODO: check length of message and break message into multiple lines
func msg(ircc *irc.IRC, w io.Writer, input string) int {
	cmd := events.Msg{}

	if err := json.Unmarshal([]byte(input), &cmd); err != nil {
		fmt.Fprint(w, events.JSONError(err.Error()))
		return -1 // Not a critical error.
	}

	e := []validate.ValidationMsg{}
	e = append(e, validate.NotNil("target", cmd.Target)...)
	e = append(e, validate.NotNil("msg", cmd.Msg)...)

	if len(e) > 0 {
		fmt.Fprint(w, events.ValidationError("join", e))
		return -1 // Not a critical error.
	}

	ircc.Msg(cmd.Target, cmd.Msg, cmd.Notice)
	return -1
}

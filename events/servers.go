package events

import "encoding/json"

// User is the format for the user command.
type User struct {
	User string `json:"user"`
	Name string `json:"name"`
}

// Nick is the format for the nick command.
type Nick struct {
	Nick string `json:"nick"`
}

// Quit is the format for the quit event
type Quit struct {
	Type   string `json:"type"`
	Status string `json:"status"`
	User   string `json:"user"`
	Msg    string `json:"msg"`
}

// Connected returns a connection event
func Connected(server, msg string) string {
	event, err := json.Marshal(StatusTargetMsgEvent{Type: "connected",
		Status: "ok",
		Target: server,
		Msg:    msg,
	})

	if err != nil {
		return InternalError(err.Error())
	}

	return string(event)
}

// RcvedQuit returns a quit event
func RcvedQuit(user, msg string) string {
	event, err := json.Marshal(Quit{Type: "quit",
		Status: "ok",
		User:   user,
		Msg:    msg,
	})

	if err != nil {
		return InternalError(err.Error())
	}

	return string(event)
}

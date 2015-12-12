package events

import "encoding/json"

// Connected returns a connection event
func Connected(server, msg string) string {
	event, err := json.Marshal(StatusTargetMsgEvent{Type: "connection",
		Status: "ok",
		Target: server,
		Msg:    msg,
	})

	if err != nil {
		return InternalError(err.Error())
	}

	return string(event)
}

// Quit returns a quit event
func Quit(server string) string {
	event, err := json.Marshal(StatusMsgEvent{Type: "quit",
		Status: "ok",
		Msg:    server,
	})

	if err != nil {
		return InternalError(err.Error())
	}

	return string(event)
}

// WaitForConnection returns a connection in progress message
func WaitForConnection(time float64) string {
	event, err := json.Marshal(StatusMsgDurationEvent{Type: "connection",
		Status:   "working",
		Msg:      "waiting for connection",
		Duration: time,
	})

	if err != nil {
		return InternalError(err.Error())
	}

	return string(event)
}

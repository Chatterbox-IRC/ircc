package irc

import (
	"fmt"

	"github.com/chatterbox-irc/chatterbox/ircc/events"
)

// User sets user and real name.
func (i IRC) User(user, name string) {
	if err := i.Write(userMsg(user, name)); err != nil {
		fmt.Fprint(i.Output, events.WriteError(err.Error()))
	}
}

// Nick sets nickname.
func (i IRC) Nick(nick string) {
	if err := i.Write(nickMsg(nick)); err != nil {
		fmt.Fprint(i.Output, events.WriteError(err.Error()))
	}
}

// Join an IRC channel.
func (i IRC) Join(channel, password string) {
	if err := i.Write(joinMsg(channel, password)); err != nil {
		fmt.Fprint(i.Output, events.WriteError(err.Error()))
	}
}

// Part from a channel.
func (i *IRC) Part(channel string) {
	if err := i.Write(partMsg(channel)); err != nil {
		fmt.Fprint(i.Output, events.WriteError(err.Error()))
	}
}

// Quit from an IRC server. You probably want IRC.Disconnect instead, which cleans up.
func (i *IRC) Quit() {
	if err := i.Write(quitMsg("disconnecting")); err != nil {
		fmt.Fprint(i.Output, events.WriteError(err.Error()))
	}
}

// Msg sends a message to a user or channel
func (i *IRC) Msg(target, msg string, notice bool) {
	if notice {
		if err := i.Write(noticeMsg(target, msg)); err != nil {
			fmt.Fprint(i.Output, events.WriteError(err.Error()))
		}
		return
	}

	if err := i.Write(privMsg(target, msg)); err != nil {
		fmt.Fprint(i.Output, events.WriteError(err.Error()))
	}
}

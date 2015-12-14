package tests

import (
	"testing"

	"github.com/chatterbox-irc/ircc/mock"
)

func TestConnection(t *testing.T) {
	t.Parallel()

	out, ircc, ircd, err := mock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer ircd.Close()

	err = ircc.Connect()
	if err != nil {
		t.Fatal(err)
	}

	ircc.User("test", "Doctor Who")
	ircc.Nick("test")

	err = mock.PollForEvent(`{"type":"connection","status":"ok","target":"hostname.domain.tld","msg":"Welcome to the ShadowNET Internet Relay Chat Network test"}`, out)
	if err != nil {
		t.Error(err)
	}
}

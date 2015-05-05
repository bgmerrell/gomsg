package messages

import (
	"testing"
)

func TestMessage(t *testing.T) {
	const (
		sender    = "alice"
		receiver  = "bob"
		testMsg   = "test message"
		nMsgsRead = 1
	)
	mm := NewMessageMap()
	mm.Post(sender, receiver, testMsg)
	msgs, err := mm.Read(receiver, nMsgsRead)
	if err != nil {
		t.Fatal(err)
	}
	if len(msgs) != nMsgsRead {
		t.Fatalf("len(msgs) = %d, want: %d", len(msgs), nMsgsRead)
	}
	if msgs[0].Body != testMsg {
		t.Errorf("msg body = %s, want: %s", msgs[0].Body, testMsg)
	}
	if msgs[0].From != sender {
		t.Errorf("msg from = %s, want: %s", msgs[0].From, sender)
	}
}

func TestMessageTruncate(t *testing.T) {
	const (
		sender    = "neil"
		receiver  = "houston"
		nMsgsRead = 2
	)
	testMsgs := []string{
		"OK, I'm going to step off the LEM now",
		"That's one small step for man",
		"One giant leap for makind"}
	mm := NewMessageMap()
	mm.Post(sender, receiver, testMsgs[0])
	mm.Post(sender, receiver, testMsgs[1])
	mm.Post(sender, receiver, testMsgs[2])
	msgs, err := mm.Read(receiver, nMsgsRead)
	if err != nil {
		t.Fatal(err)
	}
	if len(msgs) != nMsgsRead {
		t.Fatalf("len(msgs) = %d, want: %d", len(msgs), nMsgsRead)
	}
	for i := 0; i < nMsgsRead; i++ {
		if msgs[i].Body != testMsgs[i] {
			t.Errorf("msg body = %s, want: %s", msgs[i].Body, testMsgs[i])
		}
		if msgs[i].From != sender {
			t.Errorf("msg from = %s, want: %s", msgs[i].From, sender)
		}
	}
}

func TestNoMessage(t *testing.T) {
	const (
		nMsgsRead     = 10
		nMsgsExpected = 0
	)
	mm := NewMessageMap()
	msgs, err := mm.Read("bob", nMsgsRead)
	if err != nil {
		t.Fatal(err)
	}
	if len(msgs) != nMsgsExpected {
		t.Fatalf("Want empty slice")
	}
}

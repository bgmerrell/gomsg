package web

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bgmerrell/gomsg/messages"
)

func TestPost(t *testing.T) {
	const (
		testMsg  = "testing 1 2 3"
		nMsgs    = 1
		srcUser  = "alice"
		destUser = "bob"
	)
	w := httptest.NewRecorder()
	postBody := bytes.NewBufferString(testMsg)
	r, err := http.NewRequest(
		"POST",
		fmt.Sprintf("http://example.com:8080/message?from=%s&to=%s", srcUser, destUser),
		postBody)
	if err != nil {
		t.Fatal(err)
	}
	mm := messages.NewMessageMap()
	h := &Handler{mm, 0}
	h.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("w.Code = %d, want: %d", w.Code, http.StatusOK)
	}
	if w.Body.String() != testMsg {
		t.Errorf("w.Body.String() = %s, want: %s", w.Body.String(), testMsg)
	}

	// Now check to make sure the message made it into the message map
	msgs, err := mm.Read(destUser, nMsgs)
	if len(msgs) != nMsgs {
		t.Fatalf("len(msgs) = %d, want: %d", len(msgs), nMsgs)
	}
	if msgs[0].Body != testMsg {
		t.Errorf("message = %s, want: %s", msgs[0], testMsg)
	}

	// Try reading a message for an invalid user
	invalidDestUser := "bob loblaw"
	msgs, err = mm.Read(invalidDestUser, nMsgs)
	if len(msgs) != 0 {
		t.Errorf("got %d unexpected message(s) for invalid user: %#v", msgs)
	}
}

// TestHead tests the unsupported HTTP method: "HEAD"
func TestHead(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("HEAD", "http://example.com:8080/message", nil)
	if err != nil {
		t.Fatal(err)
	}
	mm := messages.NewMessageMap()
	h := &Handler{mm, 0}
	h.ServeHTTP(w, r)
	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("w.Code = %s, want: %s", w.Code, http.StatusMethodNotAllowed)
	}
	expected := "GET, POST"
	if w.Header().Get("Allow") != expected {
		t.Errorf("Allow header = %s, want: %s", w.Header().Get("Allow"), expected)
	}
}

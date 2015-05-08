package web

import (
	"io/ioutil"
	"net/http"

	"github.com/bgmerrell/gomsg/messages"
)

const (
	maxBodySize   = 1024
	toUserParam   = "to"
	fromUserParam = "from"
)

type Handler struct {
	Messenger messages.Messenger
	Count     uint64
}

type response struct {
	output string
	code   int
}

func newResponse() *response {
	return &response{"", http.StatusOK}
}

func (h *Handler) post(w http.ResponseWriter, r *http.Request) *response {
	vals := r.URL.Query()
	toUser := vals.Get(toUserParam)
	if toUser == "" {
		return &response{"no destination user specified", http.StatusBadRequest}
	}
	fromUser := vals.Get(fromUserParam)
	if fromUser == "" {
		return &response{"no source user specified", http.StatusBadRequest}
	}
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return &response{err.Error(), http.StatusInternalServerError}
	}
	err = h.Messenger.Post(fromUser, toUser, string(msg))
	if err != nil {
		return &response{err.Error(), http.StatusInternalServerError}
	}
	return &response{string(msg), http.StatusOK}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var resp *response
	r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)
	switch r.Method {
	case "POST":
		resp = h.post(w, r)
	default:
		// HTTP/1.1 spec says we must indicate which methods we allow
		w.Header().Set("Allow", "GET, POST")
		resp = &response{
			http.StatusText(http.StatusMethodNotAllowed),
			http.StatusMethodNotAllowed}
	}
	if resp.code == http.StatusOK {
		w.Write([]byte(resp.output))
	} else {
		http.Error(w, resp.output, resp.code)
	}
}

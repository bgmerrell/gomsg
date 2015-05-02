package web

import (
	"fmt"
	"html"
	"net/http"
)

type Handler struct {
	count uint64
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
}

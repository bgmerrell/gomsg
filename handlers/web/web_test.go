package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "http://example.com:8080/message", nil)
	if err != nil {
		t.Fatal(err)
	}
	h := &Handler{}
	h.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("w.Code = %d, want: %d", w.Code, http.StatusOK)
	}
	expectedBody := "Hello, \"/message\"\n"
	if w.Body.String() != expectedBody {
		t.Errorf("w.Body.String() = %s, want: %s", w.Body.String(), expectedBody)
	}
}

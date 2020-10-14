package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecureHeaders(t *testing.T) {
	w := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	secureHeaders(next).ServeHTTP(w, r)

	res := w.Result()

	frameOptions := res.Header.Get("X-Frame-Options")
	if frameOptions != "deny" {
		t.Errorf("want %q; got %q", "deny", frameOptions)
	}

	xssProtection := res.Header.Get("X-XSS-Protection")
	if xssProtection != "1; mode=block" {
		t.Errorf("want %q; got %q", "1; mode=block", xssProtection)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, res.StatusCode)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}
}

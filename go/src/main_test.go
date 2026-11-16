package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func assertJSON(t *testing.T, rec *httptest.ResponseRecorder, key, want string) {
	t.Helper()
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	var body map[string]string
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	if body[key] != want {
		t.Errorf("%s = %q, want %q", key, body[key], want)
	}
}

func TestHello(t *testing.T) {
	rec := httptest.NewRecorder()
	hello(rec, httptest.NewRequest(http.MethodGet, "/hello", nil))
	assertJSON(t, rec, "message", "Hello, World!")
}

func TestHealthy(t *testing.T) {
	rec := httptest.NewRecorder()
	healthy(rec, httptest.NewRequest(http.MethodGet, "/healthy", nil))
	assertJSON(t, rec, "status", "healthy")
}

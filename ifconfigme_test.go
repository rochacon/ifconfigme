package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIPfromXForwardedFor(t *testing.T) {
	req, err := http.NewRequest("GET", "http://ifconfigme/", nil)
	if err != nil {
		t.Error(err)
	}
	req.Header["X-Forwarded-For"] = []string{"127.0.0.2"}

	rec := httptest.NewRecorder()

	ifconfig(rec, req)

	expected := "127.0.0.2"

	if rec.Body.String() != expected {
		t.Errorf("Return IP unmatch: %s != %s", rec.Body.String(), expected)
	}
}

func TestIPfromRemoteAddr(t *testing.T) {
	req, err := http.NewRequest("GET", "http://ifconfigme/", nil)
	if err != nil {
		t.Error(err)
	}
	req.RemoteAddr = "127.0.0.1:8000"

	rec := httptest.NewRecorder()

	ifconfig(rec, req)

	expected := "127.0.0.1"

	if rec.Body.String() != expected {
		t.Errorf("Return IP unmatch: %s != %s", rec.Body.String(), expected)
	}
}

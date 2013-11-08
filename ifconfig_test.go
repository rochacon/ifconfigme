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

	if rec.Body.String() != req.Header["X-Forwarded-For"][0] {
		t.Errorf("Return IP unmatch: %s != %s", rec.Body.String(), req.Header["X-Forwarded-For"][0])
	}
}

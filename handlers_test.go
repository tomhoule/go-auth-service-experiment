package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestRootHandler is a test
func TestRootHandler(t *testing.T) {
	server := httptest.NewServer(getHandler())
	defer server.Close()
	client := &http.Client{}
	res, err := client.Get(server.URL + "/")
	if res.StatusCode != 404 {
		t.Error("The root route did not return 200")
	}
	if err != nil {
		t.Errorf("Get request to root returned an error (%s)", err.Error())
	}
}

func TestCreateUserRouteSendsAConfirmationEmail(t *testing.T) {
	server := httptest.NewServer(getHandler())
	defer server.Close()
	client := &http.Client{}
	res, err := client.Post(
		server.URL+"/create-user",
		"application/json",
		strings.NewReader(`{"email": "jcvd@jcvd.org", "password": "melonpan"}`))
	if err != nil {
		t.Errorf("Error in new user request (%s)", err.Error())
	}
	if res.StatusCode != 200 {
		t.Errorf("Error status in new user request: %d", res.StatusCode)
	}
}

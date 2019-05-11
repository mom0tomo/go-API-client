package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var sampleHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HTTP Test")
})

func TestClient(t *testing.T) {
	ts := httptest.NewServer(sampleHandler)
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error by ioutil.ReadAll(). %v", err)
	}

	if string(body) != "HTTP Test" {
		t.Fatalf("Data Error. %v", string(body))
	}
}

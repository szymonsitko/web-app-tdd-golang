package webapp_test

import (
	"log"
	"net/http"
	"testing"
)

func TestHttpResponseAndContent(t *testing.T) {
	// Testing response on 'main' page, '/' url
	resp, err := http.Get("http://localhost:8000")
	if err != nil {
		log.Fatalf("Cannot read response from server %s", err)
	} else {
		log.Printf("Status: %s", resp.Status)
	}
}

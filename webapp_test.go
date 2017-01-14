package webapp_test

import (
	// "app"
	"log"
	"net/http"
	"testing"
)

func TestHttpResponseAndContent(t *testing.T) {
	if resp, err := http.Get("http://localhost:8000"); err != nil {
		log.Fatalf("Cannot read response from server %s", err)
	} else {
		log.Print("Status %s", resp)
	}
}

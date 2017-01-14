package webapp_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
	defer resp.Body.Close()
	// Get content from http client
	htmlData, _ := ioutil.ReadAll(resp.Body)
	string_response := string(htmlData)
	if strings.Contains(string_response, "Ticket booker") != true {
		log.Fatalf("Content not found in page body")
	}
}

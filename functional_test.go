package main

import (
	"log"
	"testing"

	"github.com/tebeka/selenium"
)

func TestFunctionalMainPage(t *testing.T) {
	// Initialize functional test with selenium
	var webDriver selenium.WebDriver
	var err error
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	if webDriver, err = selenium.NewRemote(caps, "http://localhost:9515"); err != nil {
		log.Fatalf("Failed to open session: %s\n", err)
	}
	defer webDriver.Quit()
	// Setup the main page with port specified
	// in this case job is done locally at the first
	err = webDriver.Get("https://localhost:8000")
	if err != nil {
		log.Fatalf("Failed to load page: %s\n", err)
	}
	// check title functionally
	if page_title, err := webDriver.Title(); err != nil {
		webDriver.Quit()
		log.Fatalf("Cannot obtain page title %s", err)
	} else if page_title != "Welcome to The Ticket Booker!" {
		webDriver.Close()
		log.Fatalf("Title not found")
	}

}

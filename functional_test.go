package main

import (
	"log"
	"testing"
	"strings"

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
	err = webDriver.Get("http://localhost:8000")
	if err != nil {
		log.Fatalf("Failed to load page: %s\n", err)
	}
	// check title functionally
	// Searching for: title
	page_title, _ := webDriver.Title()
	if page_title != "Welcome to The Ticket Booker!" {
		webDriver.Close()
		log.Fatalf("Title not found")
	}
	// Test on content passed into document body with template
	// Searching for: h1 tag
	var elem selenium.WebElement
    elem, _ = webDriver.FindElement(selenium.ByTagName, "h1")
    header, _ := elem.Text()
    if strings.Contains(header, "Log-in") != true {
    	webDriver.Close()
    	log.Fatalf("Context not found in document body")
    }
}

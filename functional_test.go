package main

import (
	"log"
	"strings"
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
	err = webDriver.Get("http://localhost:8000")
	if err != nil {
		log.Fatalf("Failed to load page: %s\n", err)
	}
	var elem selenium.WebElement
	elem, _ = webDriver.FindElement(selenium.ByID, "register-page")
	elem.Click()
	// check title functionally
	// Searching for: title
	page_title, _ := webDriver.Title()
	if page_title != "Welcome to The Ticket Booker!" {
		webDriver.Close()
		log.Fatalf("Title not found")
	}
	// Test on content passed into document body with template
	// Searching for: h1 tag
	elem, _ = webDriver.FindElement(selenium.ByTagName, "h2")
	header, _ := elem.Text()
	if strings.Contains(header, "Log-in") != true {
		webDriver.Close()
		log.Fatalf("Context not found in document body")
	}
	// Test user login / registration form functionally
	elem, _ = webDriver.FindElement(selenium.ByID, "input-box-username")
	elem.SendKeys("example")
	elem, _ = webDriver.FindElement(selenium.ByID, "input-box-password")
	elem.SendKeys("mypassword")
	elem, _ = webDriver.FindElement(selenium.ByID, "input-box-email")
	elem.SendKeys("example@dot.com")
	elem, _ = webDriver.FindElement(selenium.ByCSSSelector, "#submit-button")
	elem.Click()
	// Check if user reached details page & bo back to the main page
	elem, _ = webDriver.FindElement(selenium.ByID, "user-details")
	content, _ := elem.Text()
	if strings.Contains(content, "User bookings details for User: example") != true {
		webDriver.Close()
		log.Fatalf("Content for user details page not found.")
	}
	// send user back to main page (with the choice prompt)
	elem, _ = webDriver.FindElement(selenium.ByTagName, "a")
	elem.Click()
}

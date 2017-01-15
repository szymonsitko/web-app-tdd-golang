package webapp_test

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func GetHttp(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Cannot read response from server %s", err)
	}
	return resp
}

var http_response = GetHttp("http://localhost:8000")

func TestHttpResponse(t *testing.T) {
	if http_response.Status == "200 OK" {
		log.Printf("Status: %s\n", http_response.Status)
	} else {
		log.Fatalf("Response status: ", http_response.Status)
	}
}

func TestContentFromResponse(t *testing.T) {
	htmlData, _ := ioutil.ReadAll(http_response.Body)
	string_response := string(htmlData)
	if strings.Contains(string_response, "Ticket booker") != true {
		log.Fatalf("Content not found in page body")
	}
}

func TestDatabaseConnectionAndExecution(t *testing.T) {
	db, err := sql.Open("mysql", "simon:irekdudek@tcp/web_go")
	if err != nil {
		log.Fatalf("Error db connection: %s", err)
	}
	defer db.Close()
}

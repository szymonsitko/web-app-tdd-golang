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

var http_response = GetHttp("http://localhost:8000/details")

func TestHttpResponse(t *testing.T) {
	if http_response.Status == "200 OK" {
		log.Printf("Status: %s\n", http_response.Status)
	} else {
		log.Fatalf("Response status: %s", http_response.Status)
	}
}

func TestDatabaseInsertion(t *testing.T) {
	db, err := sql.Open("mysql", "simon:irekdudek@tcp/web_go")
	if err != nil {
		log.Fatalf("Error db connection: %s", err)
	}
	// Database check for results
	var db_user string
	_ = db.QueryRow("SELECT username FROM users WHERE email=?;", "example@dot.com").Scan(&db_user)
	if db_user != "example" {
		log.Print("User not inserted into database\n")
	}
	// check for http response from details subpage
	htmlData, _ := ioutil.ReadAll(http_response.Body)
	string_response := string(htmlData)
	if strings.Contains(string_response, "User: example") != true {
		log.Print("Content not found in page body")
	}
	// Database cleanup
	_, err = db.Exec("DELETE FROM users;")
	err = db.Close()
	if err != nil {
		log.Println("Error closing database %s", err)
	}
}

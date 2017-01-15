package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"log"
	"testing"
)

type testUserDetails struct {
	Username string
	Password string
	Email    string
}

func TestPasswordEncrypt(t *testing.T) {
	db, _ := sql.Open("mysql", "simon:irekdudek@tcp/web_go")
	defer db.Close()
	// Test everything!
	user := testUserDetails{
		Username: "example",
		Password: "mypassword",
		Email:    "example@dot.com",
	}
	// // No error, but do I get right salted password? Let's check
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error encountered: ", err)
	}
	// Let's add real life example with the database data
	var db_password string
	// All sql operations to perform final comparison
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, username VARCHAR(255) UNIQUE, password VARCHAR(500), email VARCHAR(355) UNIQUE);")
	_, err = db.Exec("INSERT INTO users(username, password, email) VALUES(?, ?, ?)", user.Username, string(hashedPassword), user.Email)
	err = db.QueryRow("SELECT password FROM users WHERE username=?;", user.Username).Scan(&db_password)
	// Compare passwords salted vs string
	err = bcrypt.CompareHashAndPassword([]byte(db_password), []byte(user.Password))
	if err != nil {
		log.Printf("Comparison failed: %s", err)
	}
	_, err = db.Exec("DROP TABLE users;")
	// Worked!
}

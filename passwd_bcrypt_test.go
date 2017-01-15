package main

import (
	"testing"
	"log"
	"golang.org/x/crypto/bcrypt"
)

type testUserDetails struct {
	Username string
	Password string
}

func TestPasswordEncrypt(t *testing.T) {
	// Test everything!
	user := testUserDetails{Username: "example", Password: "mypassword",}
	// // No error, but do I get right salted password? Let's check
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error encountered: ", err)
	}
	// Compare passwords salted vs string
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(user.Password))
	if err != nil {
		log.Fatalf("Comparison failed")
	}
	// Worked!
}
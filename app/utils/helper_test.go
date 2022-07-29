package utils

import (
	"log"
	"testing"
)

// func TestJewetw(t *testing.T) {
// 	encryp, err := EncryptToken("taik.com", 5)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	log.Print(encryp)

// 	tk := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTkwNDE4NzUsInVzZXJfZW1haWwiOiJ0YWlrLmNvbSJ9.y9NrCMr7ehrdEoq2oyQAr4jc7k17fzVstph5WVYCPCY"

// 	data, ok := DecryptToken(tk)

// 	if !ok {
// 		t.Fatal("expaired")
// 	}

// 	log.Print(data)
// }

func TestFileName(t *testing.T) {
	file := CheckFile("user_login.json")

	if file != nil {
		t.Fatal(file.Error())
	}

	log.Print("file exist")
}

package main

import (
	"fmt"
	"golte/app"
	"log"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"testing"
)

func TestGetFilename(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	t.Logf("Current test filename: %s", filename)
}

func TestGetRequestWithoutLogin(t *testing.T) {
	router := app.SetupRouter()
	ts := httptest.NewServer(router)
	defer ts.Close()

	log.Print("url : " + ts.URL)
	resp, err := http.Get(fmt.Sprintf("%s/", ts.URL))

	if err != nil {
		t.Fatalf("you got error : %s", err.Error())
	}

	if resp.StatusCode != 200 {
		t.Fatalf("you got error in status code")
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("you got error in content-type")
	}

	if strings.Split(val[0], ";")[0] != "text/html" {
		t.Fatalf("missing content-type")
	}
}

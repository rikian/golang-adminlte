package config

import "testing"

func TestConnectionPostgres(t *testing.T) {
	err := ConnectDB()
	if err != nil {
		t.Fatal(err.Error())
		return
	}
}

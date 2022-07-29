package repository

import (
	"testing"

	"golte/app/config"
)

func TestMigrations(t *testing.T) {
	err := config.ConnectDB()
	if err != nil {
		t.Fatal(err.Error())
	}
	db := config.DB

	err = RunMigration(db)

	if err != nil {
		t.Fatal(err.Error())
	}
}

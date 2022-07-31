package repository

import (
	"log"
	"testing"

	"golte/app/config"
	user_entity "golte/app/entity/user"
)

func TestMigrations(t *testing.T) {
	err := config.ConnectDB()
	if err != nil {
		t.Fatal(err.Error())
	}
	db := config.DB

	result := db.Create(&user_entity.StatusUser{
		Id:     1,
		Status: "admin",
	})

	if result.Error != nil {
		log.Print(result.Error.Error())
	}
}

// func TestMigrations(t *testing.T) {
// 	err := config.ConnectDB()
// 	if err != nil {
// 		t.Fatal(err.Error())
// 	}
// 	db := config.DB

// 	err = RunMigration(db)

// 	if err != nil {
// 		t.Fatal(err.Error())
// 	}
// }

package user

import (
	"golte/app/config"
	"log"
	"testing"
)

var user_repo = New()

// func TestReadAllUser(t *testing.T) {
// 	_ = config.ConnectDB()
// 	user_repo.ReadAllUser()
// 	log.Print(user_repo)
// }

func TestGetTokenUser(t *testing.T) {
	_ = config.ConnectDB()

	token := user_repo.GetTokenUserById("a3")
	if token == nil {
		t.Fatal("token not found")
	}

	log.Print(token.SessionUser)
}

// func TestUpdateOneUser(t *testing.T) {
// 	_ = config.ConnectDB()
// 	request := &user_entity.UserJwt{
// 		UserEmail: "rikianfaisal@gmail.com",
// 		TokenJwt:  "12345",
// 	}
// 	data := user_repo.InsertJwtTokwn(request)

// 	if data != nil {
// 		t.Fatal(data.Error())
// 	}

// 	log.Print(data)
// }

// func TestReadOneUser(t *testing.T) {
// 	_ = config.ConnectDB()
// 	request := &user_entity.UserRequestLogin{
// 		UserEmail:    "frizka@gmail.com",
// 		UserPassword: "12345",
// 	}
// 	data, err := user_repo.ReadOneUser(request)

// 	if err != nil {
// 		t.Fatal(err.Error())
// 	}

// 	log.Print(data)
// }

// func TestPostUser(t *testing.T) {
// 	_ = config.ConnectDB()
// 	user := &user_entity.User{
// 		IdUser:       "a3",
// 		NameUser:     "boim",
// 		EmailUser:    "boim@gmail.com",
// 		PasswordUser: "12345",
// 		StatusUserId: 1,
// 		SessionUser:  "12345",
// 		CreateDate:   "7 apr",
// 		LastUpdate:   "",
// 	}

// 	err := user_repo.PostUser(user)

// 	if err != nil {
// 		t.Fatal(err.Error())
// 	}
// }

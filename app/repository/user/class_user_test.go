package user

import (
	"golte/app/config"
	"log"
	"testing"
)

func TestClassUser(t *testing.T) {
	_ = config.ConnectDB()

	user := InitUser()
	log.Print(user.ReadAllUser())
}

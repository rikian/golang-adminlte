package user

import (
	"golte/app/config"
	user_entity "golte/app/entity/user"
	"log"
)

func InitUser() *classUser {
	return &classUser{}
}

type classUser struct {
	user  *user_entity.User
	users *[]user_entity.User
}

func (cu *classUser) ReadAllUser() *[]user_entity.User {
	db := config.DB
	err := db.Model(&cu.user).Preload("StatusUser").Find(&cu.users).Error

	if err != nil {
		log.Print(err.Error())
		return nil
	}

	return cu.users
}

func (cu *classUser) ReadOneUser() *user_entity.User {
	db := config.DB
	err := db.Model(&cu.user).Preload("StatusUser").Find(&cu.user).Error

	if err != nil {
		log.Print(err.Error())
		return &user_entity.User{}
	}

	return cu.user
}

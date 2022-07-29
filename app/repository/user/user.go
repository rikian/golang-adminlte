package user

import (
	"golte/app/config"
	user_entity "golte/app/entity/user"
	"log"
)

type UserContract interface {
	ReadAllUser() []user_entity.User
	ReadOneUser(request *user_entity.UserRequestLogin) (user_entity.User, error)
	GetTokenUserById(id string) *user_entity.UserSession
	PostUser(user *user_entity.User) error
	InsertJwtTokwn(dataJwt *user_entity.UserJwt) error
}

type user struct {
	userOne user_entity.User
	user    []user_entity.User
}

func New() UserContract {
	return &user{}
}

func (use *user) ReadAllUser() []user_entity.User {
	db := config.DB
	err := db.Model(&user_entity.User{}).Preload("StatusUser").Find(&use.user).Error

	if err != nil {
		log.Print(err.Error())
		return nil
	}

	return use.user
}

func (use *user) ReadOneUser(request *user_entity.UserRequestLogin) (user_entity.User, error) {
	db := config.DB
	err := db.Model(&user_entity.User{}).Preload("StatusUser").Where("email_user = ? AND password_user = ?", request.UserEmail, request.UserPassword).Find(&use.userOne).Error

	if err != nil {
		log.Print(err.Error())
		return user_entity.User{}, err
	}

	return use.userOne, nil
}

func (u *user) PostUser(user *user_entity.User) error {
	db := config.DB
	result := db.Create(user)

	if result.Error != nil {
		log.Print(result.Error.Error())
		return result.Error
	}

	log.Print(result.RowsAffected)
	return nil
}

func (u *user) InsertJwtTokwn(dataJwt *user_entity.UserJwt) error {
	db := config.DB
	result := db.Model(&user_entity.User{}).Where("email_user = ?", dataJwt.UserEmail).Update("session_user", dataJwt.TokenJwt).Error

	if result != nil {
		return result
	}

	return nil
}

func (u *user) GetTokenUserById(id string) *user_entity.UserSession {
	var userSession *user_entity.UserSession
	db := config.DB
	db.Raw("SELECT users.session_user FROM users where id_user = ?", id).Scan(&userSession)

	return userSession
}

package controllers

import (
	"encoding/json"
	user_entity "golte/app/entity/user"
	"golte/app/repository/user"
	"golte/app/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostLogin(ctx *gin.Context) {
	var user_repo = user.New()
	var userRequestLogin *user_entity.UserRequestLogin
	err := json.NewDecoder(ctx.Request.Body).Decode(&userRequestLogin)
	if err != nil {
		ctx.JSON(200, gin.H{
			"method": "auth",
			"status": "failed",
		})
		return
	}

	if !utils.IsValidEmail(userRequestLogin.UserEmail) {
		ctx.JSON(200, gin.H{
			"method": "auth",
			"status": "failed",
		})
		return
	}

	data, err := user_repo.ReadOneUser(&user_entity.UserRequestLogin{
		UserEmail:    userRequestLogin.UserEmail,
		UserPassword: utils.SHA256(userRequestLogin.UserPassword),
	})

	if err != nil {
		ctx.JSON(200, gin.H{
			"method": "auth",
			"status": "failed",
		})
		return
	}

	if data.EmailUser == "" {
		ctx.JSON(200, gin.H{
			"method": "auth",
			"status": "failed",
		})
		return
	}

	cookie, err := utils.EncryptToken(data.EmailUser, 18000)

	if err != nil {
		ctx.JSON(200, gin.H{
			"method": "auth",
			"status": "failed",
		})
		return
	}

	// push to db
	err = user_repo.InsertJwtTokwn(&user_entity.UserJwt{
		UserEmail: data.EmailUser,
		TokenJwt:  cookie,
	})

	if err != nil {
		ctx.JSON(200, gin.H{
			"method": "auth",
			"status": "failed",
		})
		return
	}

	ctx.SetCookie("id", data.IdUser, 1800, "/", "", true, true)
	ctx.SetCookie("user", strings.Split(cookie, ".")[2], 1800, "/", "", true, true)
	ctx.JSON(200, gin.H{
		"method": "auth",
		"status": "ok",
	})
}

func PostRegister(ctx *gin.Context) {
	var user_repo = user.New()
	var userRegister user_entity.UserRegister
	err := json.NewDecoder(ctx.Request.Body).Decode(&userRegister)
	if err != nil {
		ctx.JSON(200, gin.H{
			"method": "register",
			"status": "failed",
		})

		return
	}

	if !utils.IsValidEmail(userRegister.UserEmail) {
		ctx.JSON(200, gin.H{
			"method": "register",
			"status": "failed",
		})
		return
	}

	// check password
	if userRegister.UserPassword1 != userRegister.UserPassword2 {
		ctx.JSON(200, gin.H{
			"method": "register",
			"status": "failed",
		})
		return
	}

	// check agrement
	if !userRegister.UserTerm {
		ctx.JSON(200, gin.H{
			"method": "register",
			"status": "failed",
		})
		return
	}

	register := user_repo.PostUser(&user_entity.User{
		IdUser:       uuid.NewString(),
		NameUser:     userRegister.UserName,
		EmailUser:    userRegister.UserEmail,
		PasswordUser: utils.SHA256(userRegister.UserPassword1),
		StatusUserId: 1,
		SessionUser:  "-",
		CreateDate:   time.Now().Format("01-02-2006 15:04:05"),
	})

	if register != nil {
		ctx.JSON(200, gin.H{
			"method": "register",
			"status": "failed",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"method": "register",
		"status": "ok",
	})
}

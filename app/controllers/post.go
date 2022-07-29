package controllers

import (
	"encoding/json"
	user_entity "golte/app/entity/user"
	"golte/app/repository/user"
	"golte/app/utils"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
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

	data, err := user_repo.ReadOneUser(userRequestLogin)

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
	var userRegister user_entity.UserRegister
	err := json.NewDecoder(ctx.Request.Body).Decode(&userRegister)
	if err != nil {
		ctx.JSON(200, gin.H{
			"method": "register",
			"status": "failed",
		})

		return
	}

	log.Print(userRegister)

	ctx.JSON(200, gin.H{
		"method": "register",
		"status": "ok",
	})
}

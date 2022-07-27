package controllers

import (
	"encoding/json"
	"golte/app/entity"

	"github.com/gin-gonic/gin"
)

func Post(ctx *gin.Context) {
	var user entity.User
	err := json.NewDecoder(ctx.Request.Body).Decode(&user)
	if err != nil {
		ctx.JSON(200, gin.H{
			"method": "login",
			"status": "failed",
		})
		return
	}

	ctx.SetCookie("users", "frizka", 60, "/", "", true, true)
	ctx.JSON(200, gin.H{
		"method": "login",
		"status": "ok",
	})
}

package controllers

import "github.com/gin-gonic/gin"

func Get(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}

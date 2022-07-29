package middlewares

import (
	"golte/app/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) {
	dataSession, err := cookieParser(ctx)

	if err {
		nextWithOutLogin(ctx)
		return
	}

	session := checkSession(dataSession)

	if session {
		ctx.Next()
		return
	}

	nextWithOutLogin(ctx)
}

func nextWithOutLogin(ctx *gin.Context) {
	ctx.Abort()
	if ctx.Request.Method == http.MethodPost {
		if ctx.Request.URL.Path == "/login" {
			controllers.PostLogin(ctx)
			return
		}

		if ctx.Request.URL.Path == "/register" {
			controllers.PostRegister(ctx)
			return
		}

		// default
		ctx.JSON(200, gin.H{
			"method": "auth",
			"status": "failed",
		})
		return
	}

	ctx.HTML(200, "login.html", nil)
}

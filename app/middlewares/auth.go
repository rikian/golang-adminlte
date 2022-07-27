package middlewares

import (
	"golte/app/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) {
	cookieUser, err := ctx.Request.Cookie("users")
	if err != nil {
		ctx.Abort()

		if ctx.Request.Method == http.MethodPost && ctx.Request.URL.Path == "/login" {
			controllers.Post(ctx)
			return
		}

		ctx.HTML(200, "login.html", nil)
		return
	}

	if cookieUser.Value == "frizka" {
		ctx.Next()
		return
	}

	ctx.Abort()
	ctx.Writer.WriteHeader(http.StatusInternalServerError)
	ctx.Writer.WriteString("Internal Server Lagi Error")
}

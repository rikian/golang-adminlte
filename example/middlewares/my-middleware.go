package middlewares

import (
	"github.com/gin-gonic/gin"
)

func MyMiddleware(ctx *gin.Context) {
	cookieUser, err := ctx.Request.Cookie("users")
	if err != nil {
		ctx.Abort()
		ctx.Writer.Write([]byte("unauthorize"))
		return
	}

	if cookieUser.Value == "frizka" {
		ctx.Next()
		return
	}

	ctx.Abort()
	ctx.Writer.WriteString("unauthorize")
}

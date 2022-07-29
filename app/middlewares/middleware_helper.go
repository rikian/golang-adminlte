package middlewares

import (
	"golte/app/repository/user"
	"golte/app/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type userData struct {
	id       string
	jwtToken string
}

func cookieParser(ctx *gin.Context) (*userData, bool) {
	cookieUser, err := ctx.Request.Cookie("user")
	if err != nil {
		return &userData{}, true
	}

	id, err := ctx.Request.Cookie("id")

	if err != nil {
		return &userData{}, true
	}

	if cookieUser.Value == "" || id.Value == "" {
		return &userData{}, true
	}

	return &userData{
		id:       id.Value,
		jwtToken: cookieUser.Value,
	}, false
}

func checkSession(data *userData) bool {
	var user_repo = user.New()
	token := user_repo.GetTokenUserById(data.id)

	if token == nil {
		return false
	}

	tokenParts := strings.Split(token.SessionUser, ".")

	if len(tokenParts) != 3 {
		return false
	}

	if tokenParts[2] != data.jwtToken {
		return false
	}

	decryptToken, _ := utils.DecryptToken(token.SessionUser)

	return decryptToken != nil
}

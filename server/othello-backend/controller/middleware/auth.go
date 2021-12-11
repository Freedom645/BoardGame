package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

type UserInfo struct {
	Uid string
}

func BearerMiddleware(fireApp *firebase.App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := verifyIDToken(ctx, fireApp)

		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		userInfo := UserInfo{
			Uid: token.UID,
		}

		ctx.Set("userInfo", userInfo)

		ctx.Next()
	}
}

func verifyIDToken(ctx *gin.Context, app *firebase.App) (*auth.Token, error) {
	client, err := app.Auth(ctx)
	if err != nil {
		log.Error("error getting Auth client: %v\n", err)
		return nil, err
	}

	idToken, err := parseHeader(ctx)
	if err != nil {
		log.Error("error parse header: %v\n", err)
		return nil, err
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Error("error verifying ID token: %v\n", err)
		return nil, err
	}

	return token, nil
}

func parseHeader(ctx *gin.Context) (string, error) {
	header := ctx.Request.Header.Values("Authorization")
	if len(header) == 0 {
		return "", errors.New("the header format is invalid")
	}

	values := strings.Split(header[0], " ")
	if len(values) != 2 {
		return "", errors.New("the header format is invalid")
	}
	if values[0] != "Bearer" {
		return "", errors.New("the header format is invalid")
	}

	return values[1], nil
}

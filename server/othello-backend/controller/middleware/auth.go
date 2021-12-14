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

/* 認証情報取得用のミドルウェア */
func BearerMiddleware(fireApp *firebase.App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := verifyIDToken(ctx, fireApp)

		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
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

	idToken, err := exactToken(ctx)
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

func exactToken(ctx *gin.Context) (string, error) {
	idToken, err := parseBearerHeader(ctx)
	if err == nil {
		return idToken, err
	}

	return parseWebSocketProtocol(ctx)
}

func parseBearerHeader(ctx *gin.Context) (string, error) {
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

func parseWebSocketProtocol(ctx *gin.Context) (string, error) {
	header := ctx.Request.Header.Values("Sec-WebSocket-Protocol")
	if len(header) == 0 {
		return "", errors.New("the header format is invalid")
	}

	return header[0], nil
}

/* Contextからユーザ情報取得 */
func GetUserFromContext(ctx *gin.Context) (UserInfo, error) {
	obj, exists := ctx.Get("userInfo")
	if exists {
		return obj.(UserInfo), nil
	}

	return UserInfo{}, errors.New("userinfo does not exists")
}

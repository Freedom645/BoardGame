package main

import (
	"context"
	"fmt"
	"time"

	firebase "firebase.google.com/go"
	"github.com/Freedom645/BoardGame/controller"
	"github.com/Freedom645/BoardGame/controller/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

var corsConfig = cors.New(cors.Config{
	AllowOrigins: []string{
		"http://localhost:4200",
		"http://192.168.1.52:4200",
	},
	AllowMethods: []string{
		"POST",
		"GET",
	},
	AllowHeaders: []string{
		"Access-Control-Allow-Credentials",
		"Access-Control-Allow-Headers",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"Authorization",
	},
	AllowCredentials: true,
	MaxAge:           24 * time.Hour,
})

func newFirebaseApp() *firebase.App {
	opt := option.WithCredentialsFile("service-account-file.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(fmt.Errorf("error initializing app: %v", err))
	}
	return app
}

func main() {
	log.Info("App start.")

	router := gin.Default()
	router.Use(corsConfig)
	router.Use(middleware.BearerMiddleware(newFirebaseApp()))

	rg := router.Group("/room")
	rg.GET("", controller.HandleGetRoomList)
	rg.POST("", controller.HandleCreateRoom)
	rg.GET("/:id", controller.HandleGetRoom)
	rg.GET("/:id/ws", controller.HandleConnect)

	// Listen and server on 0.0.0.0:8080
	router.Run(":8080")
}

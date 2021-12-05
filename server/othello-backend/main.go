package main

import (
	"time"

	"github.com/Freedom645/BoardGame/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("App start.")

	router := gin.Default()
	router.Use(corsConfig)

	rg := router.Group("/room")
	rg.GET("", controller.HandleGetRoomList)
	rg.POST("", controller.HandleCreateRoom)
	rg.GET("/:id", controller.HandleGetRoom)
	rg.GET("/:id/ws", controller.HandleConnect)

	// Listen and server on 0.0.0.0:8080
	router.Run(":8080")
}

var corsConfig = cors.New(cors.Config{
	AllowOrigins: []string{
		"http://localhost:4200",
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

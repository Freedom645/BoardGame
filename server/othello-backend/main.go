package main

import (
	"github.com/Freedom645/BoardGame/controller"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("App start.")

	router := gin.Default()

	rg := router.Group("/room")
	rg.GET("", controller.HandleGetRoomList)
	rg.GET("/:id", controller.HandleGetRoom)
	rg.POST("", controller.HandleCreateRoom)
	rg.GET("/ws", controller.HandleConnect)

	// Listen and server on 0.0.0.0:8080
	router.Run(":8080")
}

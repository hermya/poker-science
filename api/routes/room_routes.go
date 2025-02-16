package routes

import (
	"github.com/gin-gonic/gin"
	"poker_science/api/handlers"
)

func RegisterRoomRoutes(router *gin.Engine) {
	roomRoutes := router.Group("/rooms")
	{
		roomRoutes.POST("/create", handlers.CreateRoom) // Create a new room
		//roomRoutes.GET("/", handlers.GetRooms)      // Fetch available rooms
		//roomRoutes.POST("/join", handlers.JoinRoom) // Join a room
	}
}

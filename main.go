package main

import (
	"github.com/gin-gonic/gin"
	"poker_science/api/routes"
)

func main() {
	router := gin.Default()
	routes.RegisterRoomRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}

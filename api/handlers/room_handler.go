package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	dto "poker_science/api/dto/room"
	roomservice "poker_science/internal/application/service/room"
)

func CreateRoom(c *gin.Context) {
	request := dto.CreateRoomRequest{}
	fmt.Println(c.Request)
	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	roomUuid := roomservice.CreateRoom(request)
	response := dto.CreateRoomResponse{
		RoomUuid: roomUuid,
	}
	c.JSON(http.StatusCreated, response)
}

//
//func GetRooms(c *gin.Context) {
//	rooms := services.GetAvailableRooms()
//	c.JSON(http.StatusOK, rooms)
//}
//
//func JoinRoom(c *gin.Context) {
//	var req struct {
//		RoomID   string `json:"room_id"`
//		PlayerID string `json:"player_id"`
//	}
//
//	if err := c.ShouldBindJSON(&req); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
//		return
//	}
//
//	err := services.JoinRoom(req.RoomID, req.PlayerID)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "Joined room successfully"})
//}

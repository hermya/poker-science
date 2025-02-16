package room

type CreateRoomRequest struct {
	RequestingEntityName string  `json:"requestingEntityName" binding:"required"`
	RequestingEntityUUID string  `json:"requestingEntityUUID" binding:"required"`
	RoomSize             int8    `json:"roomSize" binding:"required"`
	TimeOut              int8    `json:"timeOut" binding:"required"`
	EntryAmount          float64 `json:"entryAmount" binding:"required"`
	SmallBlindBet        float64 `json:"smallBlindBet" binding:"required"`
}

type CreateRoomResponse struct {
	RoomUuid string `json:"roomUuid" binding:"required"`
}

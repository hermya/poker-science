package room

import (
	"fmt"
	"poker_science/api/dto/room"
	"poker_science/db/model"
)

// These are baseline logical functions which all activities around a room

func CreateRoom(request room.CreateRoomRequest) string {
	// what are the inputs ?
	// -> nameOfEntity who creates the room (Player, Server)
	// -> roomSize -> number of people allowed
	// -> timeOut  -> timeOut for auto-fold
	// -> EntryAmount -> lowest buyIn
	// -> bet -> Small blind bet

	// In order to create a room, you have to be a player, so
	// TODO fetch player first
	player := getTemporaryPlayer(request.RequestingEntityUUID)
	var playerList []model.Player
	playerList = append(playerList, player)
	temporaryRoom := model.Room{
		HostUuid:    request.RequestingEntityUUID,
		HostName:    request.RequestingEntityName,
		RoomUuid:    "room-number-1",
		Players:     playerList,
		TableSize:   request.RoomSize,
		TimeOut:     request.TimeOut,
		BuyIn:       request.EntryAmount,
		BasePrice:   request.SmallBlindBet,
		CurrentGame: model.TableState{},
	}
	fmt.Println(temporaryRoom)
	roomUuid := temporaryRoom.RoomUuid
	return roomUuid
}

func getTemporaryPlayer(playerUUID string) model.Player {
	temporaryPlayer := model.Player{
		PlayerUuid:      playerUUID,
		PlayerName:      "admin",
		PlayerAccount:   model.Account{},     // assume empty account for now
		PlayerGameState: model.PlayerState{}, // should be nullified
	}
	return temporaryPlayer
}

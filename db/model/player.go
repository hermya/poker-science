package model

type Player struct {
	PlayerUuid      string
	PlayerName      string
	PlayerAccount   Account
	PlayerGameState PlayerState
}

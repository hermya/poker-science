package player

type Player struct {
	PlayerUuid   string
	PlayerName   string
	GlobalValues GlobalMeta
	GameValues   RoomMeta
}

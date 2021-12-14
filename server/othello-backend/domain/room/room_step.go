package room

type RoomStep string

const (
	Pending  = RoomStep("pending")
	Black    = RoomStep("black")
	White    = RoomStep("white")
	GameOver = RoomStep("gameOver")
)

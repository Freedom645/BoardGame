package room

import (
	"time"

	"github.com/Freedom645/BoardGame/game"
	"github.com/Freedom645/BoardGame/room/player"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Room struct {
	uuid    uuid.UUID
	game    *game.Game
	sockets map[*websocket.Conn]bool
	created time.Time
}

func NewRoom(id uuid.UUID) *Room {
	r := new(Room)

	r.uuid = id
	r.game = game.NewGame()
	r.sockets = make(map[*websocket.Conn]bool)
	r.created = time.Now()

	return r
}

func (r *Room) UUID() uuid.UUID {
	return r.uuid
}

func (r *Room) SetFirstPlayer(player player.Player) bool {
	return r.game.SetFirstPlayer(player)
}

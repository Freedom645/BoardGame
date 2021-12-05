package room

import (
	"sync"
	"time"

	"github.com/Freedom645/BoardGame/domain/game"
	"github.com/Freedom645/BoardGame/domain/player"
	"github.com/google/uuid"
)

type Room struct {
	uuid    uuid.UUID
	game    *game.Game
	created time.Time
	locker  sync.RWMutex
}

func NewRoom(id uuid.UUID) *Room {
	r := new(Room)

	r.uuid = id
	r.game = game.NewGame()
	r.created = time.Now()

	return r
}

func (r *Room) UUID() uuid.UUID {
	return r.uuid
}

func (r *Room) Created() time.Time {
	return r.created
}

func (r *Room) SetFirstPlayer(player player.Player) bool {
	// 排他処理
	r.locker.Lock()
	defer r.locker.Unlock()

	return r.game.SetFirstPlayer(player)
}

func (r *Room) SetSecondPlayer(player player.Player) bool {
	// 排他処理
	r.locker.Lock()
	defer r.locker.Unlock()

	return r.game.SetSecondPlayer(player)
}

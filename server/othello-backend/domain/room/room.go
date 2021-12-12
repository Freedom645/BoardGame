package room

import (
	"errors"
	"sync"
	"time"

	"github.com/Freedom645/BoardGame/domain/enum/stone_type"
	"github.com/Freedom645/BoardGame/domain/game"
	"github.com/Freedom645/BoardGame/domain/player"
	"github.com/google/uuid"
)

type Room struct {
	uuid    uuid.UUID
	Game    *game.Game
	created time.Time
	locker  sync.RWMutex
}

func NewRoom(id uuid.UUID) *Room {
	r := new(Room)

	r.uuid = id
	r.Game = game.NewGame()
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

	return r.Game.SetFirstPlayer(player)
}

func (r *Room) SetSecondPlayer(player player.Player) bool {
	// 排他処理
	r.locker.Lock()
	defer r.locker.Unlock()

	return r.Game.SetSecondPlayer(player)
}

func (r *Room) Put(p game.Point, stone stone_type.StoneType) error {
	// 排他処理
	r.locker.Lock()
	defer r.locker.Unlock()

	if r.Game.Board.TypeAt(p) != stone_type.None {
		return errors.New("already placed")
	}

	points := r.Game.Board.PutIf(p, stone)

	if len(points) == 0 {
		return errors.New("there are no stones that can be turned over")
	}

	r.Game.Board.PutOne(p, stone)
	r.Game.Board.Put(&points, stone)

	return nil
}

package room

import (
	"errors"
	"sync"
	"time"

	"github.com/Freedom645/BoardGame/domain/enum/player_type"
	"github.com/Freedom645/BoardGame/domain/enum/stone_type"
	"github.com/Freedom645/BoardGame/domain/game"
	"github.com/Freedom645/BoardGame/domain/player"
	"github.com/google/uuid"
)

type Room struct {
	uuid    uuid.UUID
	Game    *game.Game
	players []player.Player
	created time.Time
	locker  sync.RWMutex
}

/* 部屋作成 */
func NewRoom() (*Room, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	r := new(Room)
	r.uuid = id
	r.Game = game.NewGame()
	r.created = time.Now()

	return r, nil
}

func (r *Room) UUID() uuid.UUID {
	return r.uuid
}

func (r *Room) Created() time.Time {
	return r.created
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

/* 部屋参加 */
func (r *Room) AddPlayer(player player.Player, playerType player_type.PlayerType) (bool, error) {
	r.locker.Lock()
	defer r.locker.Unlock()

	switch playerType {
	case player_type.Player:
		if len(r.players) >= 2 {
			// プレイヤーが2人以上いる場合は失敗
			return false, nil
		}
		r.players = append(r.players, player)
		return true, nil
	case player_type.Spectator:
		// 観戦者は何もしない
		return true, nil
	}
	return false, errors.New("unknown type")
}

/* 部屋を抜ける  */
func (r *Room) removePlayer(player player.Player) bool {
	r.locker.Lock()
	defer r.locker.Unlock()

	for i, p := range r.players {
		if p.Id == player.Id {
			// 除去
			r.players = append(r.players[:i], r.players[i+1:]...)
			break
		}
	}

	return len(r.players) == 0
}

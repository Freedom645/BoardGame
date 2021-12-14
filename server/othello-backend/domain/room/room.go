package room

import (
	"errors"
	"math/rand"
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
	step    RoomStep
	game    *game.Game
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
	r.step = Pending
	r.game = game.NewGame()
	r.created = time.Now()

	return r, nil
}

func (r *Room) UUID() uuid.UUID {
	r.locker.RLock()
	defer r.locker.RUnlock()

	return r.uuid
}

func (r *Room) Owner() player.Player {
	r.locker.RLock()
	defer r.locker.RUnlock()

	if len(r.players) == 0 {
		return player.Player{}
	}
	return r.players[0]
}

func (r *Room) Players() []player.Player {
	r.locker.RLock()
	defer r.locker.RUnlock()

	return r.players
}

func (r *Room) Step() RoomStep {
	r.locker.RLock()
	defer r.locker.RUnlock()

	return r.step
}

func (r *Room) Turn() game.Turn {
	r.locker.RLock()
	defer r.locker.RUnlock()

	return *r.game.Turn
}

func (r *Room) Stones() [][]stone_type.StoneType {
	r.locker.RLock()
	defer r.locker.RUnlock()

	return r.game.Board.Stones()
}

func (r *Room) Created() time.Time {
	r.locker.RLock()
	defer r.locker.RUnlock()

	return r.created
}

/* 承認処理 */
func (r *Room) Approve(uid string, name string, isApprove bool) error {
	r.locker.Lock()
	defer r.locker.Unlock()

	if len(r.players) != 2 {
		return errors.New("num of player is invalid")
	}

	var isApproveAll = true
	for _, v := range r.players {
		if uid == v.Id {
			v.IsApprove = isApprove
		}
		isApproveAll = isApproveAll && v.IsApprove
	}

	if isApproveAll {
		// 全員承認
		r.step = Black
		if rand.Float64() <= 0.5 {
			r.game.SetTurn(&r.players[0], &r.players[1])
		} else {
			r.game.SetTurn(&r.players[1], &r.players[0])
		}
		// 承認状態をリセット
		for _, v := range r.players {
			v.IsApprove = false
		}
	}
	return nil
}

/* 碁石を打つ */
func (r *Room) Put(uid string, p game.Point) error {
	r.locker.Lock()
	defer r.locker.Unlock()

	if r.step != Black && r.step != White {
		return errors.New("state missmatch")
	}

	stone, err := r.game.ToStone(uid)
	if err != nil {
		return err
	}

	if stone != stone_type.StoneType(r.step) {
		return errors.New("not your turn")
	}

	if r.game.Board.TypeAt(p) != stone_type.None {
		return errors.New("already placed")
	}

	points := r.game.Board.PutIf(p, stone)

	if len(points) == 0 {
		return errors.New("there are no stones that can be turned over")
	}

	r.game.Board.PutOne(p, stone)
	r.game.Board.Put(&points, stone)

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

/* 部屋を抜ける */
func (r *Room) RemovePlayer(uid string) int {
	r.locker.Lock()
	defer r.locker.Unlock()

	for i, p := range r.players {
		if p.Id == uid {
			// 順列を維持して除去
			r.players = append(r.players[:i], r.players[i+1:]...)
			break
		}
	}

	return len(r.players)
}

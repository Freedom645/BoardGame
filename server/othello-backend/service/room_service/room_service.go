package room_service

import (
	pt "github.com/Freedom645/BoardGame/domain/enum/player_type"
	"github.com/Freedom645/BoardGame/domain/player"
	"github.com/Freedom645/BoardGame/domain/room"
	"github.com/google/uuid"
)

/* 部屋作成 */
func CreateRoom() (*room.Room, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	newRoom := room.NewRoom(id)

	return newRoom, nil
}

/* 部屋参加 */
func AddPlayer(room *room.Room, player player.Player, playerType pt.PlayerType) bool {
	switch playerType {
	case pt.First:
		return room.SetFirstPlayer(player)
	case pt.Second:
		return room.SetSecondPlayer(player)
	case pt.Spectator:
		// 観戦者は何もしない
		return true
	default:
		return true
	}
}

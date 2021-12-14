package model

import (
	"time"

	"github.com/Freedom645/BoardGame/domain/player"
	"github.com/Freedom645/BoardGame/domain/room"
)

type Player struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	IsApprove bool   `json:"isApprove"`
}

type Room struct {
	Id      string    `json:"id"`
	Owner   Player    `json:"owner"`
	Players []Player  `json:"players"`
	Created time.Time `json:"created"`
}

func Of(r *room.Room) *Room {
	return &Room{
		Id:      r.UUID().String(),
		Owner:   ParsePlayer(r.Owner()),
		Players: ParsePlayers(r.Players()),
		Created: r.Created(),
	}
}

func ParsePlayers(arr []player.Player) []Player {
	var res []Player
	for _, v := range arr {
		e := ParsePlayer(v)
		res = append(res, e)
	}

	return res
}

func ParsePlayer(p player.Player) Player {
	return Player{
		Id:        p.Id,
		Name:      p.Name,
		IsApprove: p.IsApprove,
	}
}

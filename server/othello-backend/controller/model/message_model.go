package model

import (
	"github.com/Freedom645/BoardGame/controller/model/state_model"
	"github.com/Freedom645/BoardGame/domain/enum/stone_type"
	"github.com/Freedom645/BoardGame/domain/game"
)

type Message struct {
	Request  RequestMessage  `json:"request"`
	Response ResponseMessage `json:"response"`
}

type RequestMessage struct {
	Pending  PendingRequest  `json:"pending"`
	Game     GameRequest     `json:"game"`
	GameOver GameOverRequest `json:"gameOver"`
}

type ResponseMessage struct {
	Step    state_model.GameState    `json:"step"`
	Board   [][]stone_type.StoneType `json:"board"`
	Owner   Player                   `json:"owner"`
	Players []Player                 `json:"players"`
	Turn    Turn                     `json:"turn"`
}

/* 承認待ちのリクエスト */
type PendingRequest struct {
	IsApproved bool `json:"isApproved"`
}

/* 対局中リクエスト */
type GameRequest struct {
	Point Point `json:"point"`
}

/* 対局終了時リクエスト */
type GameOverRequest struct {
	IsContinued bool `json:"isContinued"`
}

/* 座標 */
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func TurnOf(t game.Turn) Turn {
	return Turn{
		BlackId: t.BlackId,
		WhiteId: t.WhiteId,
	}
}

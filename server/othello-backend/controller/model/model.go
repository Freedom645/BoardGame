package model

import (
	gc "github.com/Freedom645/BoardGame/controller/model/game_color_model"
	gs "github.com/Freedom645/BoardGame/controller/model/game_state_model"
	"github.com/Freedom645/BoardGame/domain/enum/stone_type"
)

type GameMessage struct {
	Request  GameRequestMessage  `json:"request"`
	Response GameResponseMessage `json:"response"`
}

type GameRequestMessage struct {
	PlayerName string          `json:"playerName"`
	Pending    PendingRequest  `json:"pending"`
	Game       GameRequest     `json:"game"`
	GameOver   GameOverRequest `json:"gameOver"`
}

type GameResponseMessage struct {
	State gs.GameState             `json:"state"`
	Board [][]stone_type.StoneType `json:"board"`
}

/* 参加待ちのサーバ通知 */
type MatchingResponse struct {
	RivalName string `json:"rivalName"`
}

/* 承認待ちのリクエスト */
type PendingRequest struct {
	IsApproved bool `json:"isApproved"`
}

/* 承認待ちのレスポンス */
type PendingResponse struct {
	IsApproved bool         `json:"isApproved"`
	UserName   string       `json:"userName"`
	Turn       gc.GameColor `json:"turn"`
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type GameRequest struct {
	Point Point `json:"point"`
}

type GameOverRequest struct {
	IsContinued bool `json:"isContinued"`
}

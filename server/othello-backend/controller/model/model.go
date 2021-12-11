package model

import (
	gc "github.com/Freedom645/BoardGame/controller/model/game_color_model"
	gs "github.com/Freedom645/BoardGame/controller/model/game_state_model"
	"github.com/google/uuid"
)

type GameMessage struct {
	Request  GameRequestMessage  `json:"request"`
	Response GameResponseMessage `json:"response"`
}

type GameRequestMessage struct {
	UUID     uuid.UUID `json:"uuid"`
	UserName string    `json:"userName"`
}

type GameResponseMessage struct {
	State gs.GameState `json:"state"`
}

/* 参加待ちのサーバ通知 */
type MatchingResponse struct {
	RivalName string `json:"rivalName"`
}

/* 承認待ちのリクエスト */
type PendingRequest struct {
	IsApproved bool      `json:"isApproved"`
	UUID       uuid.UUID `json:"uuid"`
	UserName   string    `json:"userName"`
}

/* 承認待ちのレスポンス */
type PendingResponse struct {
	IsApproved bool         `json:"isApproved"`
	UserName   string       `json:"userName"`
	Turn       gc.GameColor `json:"turn"`
}

/* 順番待機のレスポンス */
type WatingResponse struct {
	Turn gc.GameColor `json:"turn"`
}

package model

import (
	"github.com/google/uuid"
)

type GameRequestMessage struct {
	UUID     uuid.UUID `json:"uuid"`
	UserName string    `json:"userName"`
}

type GameState int

const (
	Joining GameState = iota + 1
	Pending
)

type GameResponseMessage struct {
	State GameState
}

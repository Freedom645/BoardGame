package game

import (
	"errors"

	"github.com/Freedom645/BoardGame/domain/enum/stone_type"
	"github.com/Freedom645/BoardGame/domain/player"
)

type Turn struct {
	BlackId string
	WhiteId string
}

type Game struct {
	/* 盤面状態 */
	Board *Board
	/* 手番 */
	Turn *Turn
}

func NewGame() *Game {
	g := new(Game)

	g.Board = NewBoard()
	g.Turn = &Turn{}

	return g
}

func (g *Game) ToStone(uid string) (stone_type.StoneType, error) {
	switch uid {
	case g.Turn.BlackId:
		return stone_type.Black, nil
	case g.Turn.WhiteId:
		return stone_type.White, nil
	}

	return stone_type.None, errors.New("not participating")
}

func (g *Game) SetTurn(black *player.Player, white *player.Player) {
	g.Turn.BlackId = black.Id
	g.Turn.WhiteId = white.Id
}

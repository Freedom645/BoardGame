package game

import (
	gs "github.com/Freedom645/BoardGame/domain/enum/game_step"
	"github.com/Freedom645/BoardGame/domain/player"
)

type Game struct {
	/* 盤面状態 */
	Board *Board
	/* 先攻プレイヤー */
	First *player.Player
	/* 後攻プレイヤー */
	Second *player.Player
	/* 状態 */
	Step gs.Step
}

func NewGame() *Game {
	g := new(Game)

	g.Board = NewBoard()
	g.Step = gs.NotStarted

	return g
}

func (g *Game) SetFirstPlayer(p player.Player) bool {
	if g.First != nil {
		return false
	}

	g.First = &p
	return true
}

func (g *Game) SetSecondPlayer(p player.Player) bool {
	if g.Second != nil {
		return false
	}

	g.Second = &p
	return true
}

func (g *Game) SetStep(step gs.Step) {
	g.Step = step
}

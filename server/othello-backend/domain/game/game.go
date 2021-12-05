package game

import (
	gs "github.com/Freedom645/BoardGame/domain/enum/game_step"
	"github.com/Freedom645/BoardGame/domain/game/board"
	"github.com/Freedom645/BoardGame/domain/player"
)

type Game struct {
	/* 盤面状態 */
	b *board.Board
	/* 先攻プレイヤー */
	first *player.Player
	/* 後攻プレイヤー */
	second *player.Player
	/* 状態 */
	step gs.Step
}

func NewGame() *Game {
	g := new(Game)

	g.b = board.NewBoard()
	g.step = gs.NotStarted

	return g
}

func (g *Game) SetFirstPlayer(p player.Player) bool {
	if g.first != nil {
		return false
	}

	g.first = &p
	return true
}

func (g *Game) SetSecondPlayer(p player.Player) bool {
	if g.second != nil {
		return false
	}

	g.second = &p
	return true
}

func (g *Game) SetStep(step gs.Step) {
	g.step = step
}

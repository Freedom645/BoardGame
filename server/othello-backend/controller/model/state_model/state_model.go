package state_model

import "github.com/Freedom645/BoardGame/domain/room"

type GameState string

const (
	/* 承認待ち */
	Pending = GameState("pending")
	/* 黒の番 */
	BlackTurn = GameState("black")
	/* 白の番 */
	WhiteTurn = GameState("white")
	/* ゲーム終了 */
	GameOver = GameState("gameOver")
)

func Of(step room.RoomStep) GameState {
	switch step {
	case room.Pending:
		return Pending
	case room.Black:
		return BlackTurn
	case room.White:
		return WhiteTurn
	case room.GameOver:
		return GameOver
	}

	return ""
}

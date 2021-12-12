package gameStateModel

type GameState string

const (
	/* 参加待ち */
	Matching GameState = "matching"
	/* 承認待ち */
	Pending GameState = "pending"
	/* ゲーム順番待ち */
	Waiting GameState = "waiting"
	/* 黒の番 */
	BlackTurn GameState = "black"
	/* 白の番 */
	WhiteTurn GameState = "white"
	/* ゲーム終了 */
	GameOver GameState = "gameOver"
	/* 継続待ち */
	Continue GameState = "continue"
)

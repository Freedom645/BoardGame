/** 碁石の種類 */
export type StoneType = "Black" | "White" | "None";

/** 座標 */
export type Point = { x: number; y: number; };

/** 参加待ち */
export type Matching = "matching";
/** 承認待ち */
export type Pending = "pending";
/** ゲーム順番待ち */
export type Waiting = "waiting";
/** 自分の番 */
export type MyTurn = "myTurn";
/** 相手の番 */
export type OpponentTurn = "opponentTurn";
/** ゲーム終了 */
export type GameOver = "gameOver";
/** 継続待ち */
export type Continue = "continue";

/** ゲーム遷移状態 */
export type GameStep = Matching | Pending | Waiting | MyTurn | OpponentTurn | GameOver | Continue;

/** ゲームリクエストメッセージ */
export type GameRequestMessage = string;

/** ゲームレスポンスメッセージ */
export type GameResponseMessage = string;

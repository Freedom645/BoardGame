import { Deserializer, Serializer } from "src/app/service/web-socket.service";
import { Board, Point, Stone, StoneType } from "./game";

/** 承認待ち */
export type Pending = "pending";
/** ゲーム順番待ち */
export type Waiting = "waiting";
/** ゲーム終了 */
export type GameOver = "gameOver";

/** 定数  */
export const Step = {
  Pending: "pending",
  Black: Stone.Black,
  White: Stone.White,
  GameOver: "gameOver",
} as {
  Pending: Pending,
  Black: StoneType,
  White: StoneType,
  GameOver: GameOver,
};

/** ゲーム遷移状態 */
export type GameStep = Pending | Waiting | StoneType | GameOver;

export const serializer: Serializer<GameMessage> = (e) => JSON.stringify(e);

export const deserializer: Deserializer<GameMessage> = (e) => JSON.parse(e.data);

/** ゲームメッセージ */
export interface GameMessage {
  request?: GameRequestMessage;
  response?: GameResponseMessage;
};

/** ゲームリクエストメッセージ */
export interface GameRequestMessage {
  playerName: string;
  pending?: PendingRequest;
  game?: GameRequest;
  gameOver?: GameOverRequest;
};

export interface PendingRequest {
  isApproved: boolean;
}

export interface GameRequest {
  point: Point;
}

export interface GameOverRequest {
  isContinued: boolean;
}

/** ゲームレスポンスメッセージ */
export interface GameResponseMessage {
  /** ゲーム遷移状態 */
  step: GameStep;
  /** 盤面状態 */
  board: Board;
};


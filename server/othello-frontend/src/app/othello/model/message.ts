import { Deserializer, Serializer } from "src/app/service/web-socket.service";
import { Point, Stone, StoneType } from "./game";

/** 参加待ち */
export type Matching = "matching";
/** 承認待ち */
export type Pending = "pending";
/** ゲーム順番待ち */
export type Waiting = "waiting";
/** ゲーム終了 */
export type GameOver = "gameOver";
/** 継続待ち */
export type Continue = "continue";

/** 定数  */
export const Step = {
  Matching: "matching",
  Pending: "pending",
  Waiting: "waiting",
  Black: Stone.Black,
  White: Stone.White,
  GameOver: "gameOver",
  Continue: "continue",
} as {
  Matching: Matching,
  Pending: Pending,
  Waiting: Waiting,
  Black: StoneType,
  White: StoneType,
  GameOver: GameOver,
  Continue: Continue,
};

/** ゲーム遷移状態 */
export type GameStep = Matching | Pending | Waiting | StoneType | GameOver | Continue;

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
};


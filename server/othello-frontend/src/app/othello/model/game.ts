
/** 碁石定数 */
export const Stone = {
  Black: "black",
  White: "white",
  None: "none",
} as {
  Black: StoneType,
  White: StoneType,
  None: StoneType,
};

/** 碁石の種類 */
export type StoneType = "black" | "white" | "none";

/** 座標 */
export interface Point { x: number; y: number; };

/** 1マス情報 */
export interface Mass {
  type: StoneType;
  point: Point;
};

/** 盤面情報 */
export type Board = StoneType[][];

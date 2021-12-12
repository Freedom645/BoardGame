
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
export type Board = Mass[];

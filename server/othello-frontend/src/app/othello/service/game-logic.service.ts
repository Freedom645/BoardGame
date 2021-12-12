import { Injectable } from '@angular/core';
import { Board, Point, Stone, StoneType } from '../model/game';

@Injectable()
export class GameLogicService {

  /** 八近傍 */
  private static readonly DIR8: Point[] = [
    { x: 0, y: -1 },
    { x: 0, y: 1 },
    { x: -1, y: 0 },
    { x: 1, y: 0 },
    { x: -1, y: -1 },
    { x: 1, y: -1 },
    { x: -1, y: 1 },
    { x: 1, y: 1 },
  ];

  /** 盤面サイズ */
  public static readonly MassNum = 8;

  constructor() { }

  /** 座標から連番へ変換 */
  public convertToIndex(point: Point): number {
    return point.y * GameLogicService.MassNum + point.x;
  }

  /** 連番から座標へ変換 */
  public convertToPoint(index: number): Point {
    return { x: index % 8, y: Math.floor(index / 8) };
  }

  /** 指定された座標の碁石タイプを取得 */
  public getType(b: Board, p: Point | number): StoneType {
    const point = this.parse(p);
    return b[point.y][point.x];
  }

  /** 盤面内の判定 */
  public contain(b: Board, p: Point | number): boolean {
    const point = this.parse(p);
    return 0 <= point.x && point.x < GameLogicService.MassNum && 0 <= point.y && point.y < GameLogicService.MassNum;
  }

  /** Point型に集約 */
  private parse(p: Point | number): Point {
    if (typeof p == "number") {
      return this.convertToPoint(p);
    }
    return p;
  }

  /** ベクトルの和 */
  public addPoint(p1: Point, p2: Point): Point {
    return { x: p1.x + p2.x, y: p1.y + p2.y };
  }

  /** ベクトルのスカラー倍 */
  public productPoint(p: Point, s: number): Point {
    return { x: p.x * s, y: p.y * s };
  }

  /** 打った場合に返される箇所の座標を返す */
  public putIf(point: Point, type: StoneType, board: Board): Point[] {
    let res: Point[] = [];
    for (let dir of GameLogicService.DIR8) {
      const list = this.searchRevPoint(point, dir, type, board);
      res = res.concat(list);
    }
    return res;
  }

  /** 反対色の碁石タイプ */
  public revStone(type: StoneType): StoneType {
    switch (type) {
      case Stone.Black: return Stone.White;
      case Stone.White: return Stone.Black;
    }
    return type;
  }

  /** 一方向へのひっくり返す碁石の座標を返す */
  private searchRevPoint(point: Point, vec: Point, type: StoneType, board: Board): Point[] {
    const res: Point[] = [];
    for (let i = 1; true; i++) {
      const now = this.addPoint(point, this.productPoint(vec, i));
      if (!this.contain(board, now)) {
        break;
      }
      const nowType = this.getType(board, now);
      if (nowType == Stone.None) {
        break;
      }

      if (nowType == type) {
        // 対になる色があったら終了
        return res;
      }

      res.push(now);
    }

    return [];
  }
}

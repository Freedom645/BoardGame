import { Component, OnInit } from '@angular/core';

interface BoardMass {
  index: number,
  x: number,
  y: number,
  width: number,
  height: number,
  bgColor: string,
}

interface BoardCover extends BoardMass {
  opacity: number,
}

interface Stone {
  index: number,
  x: number,
  y: number,
  r: number,
  stoneColor: string,
  strokeColor: string,
  strokeWidth: number,
}

const SIZE = 100;

const DEFAULT_BOARD = {
  width: SIZE - 2,
  height: SIZE - 2,
  bgColor: "green"
};

@Component({
  selector: 'app-game-board',
  templateUrl: './game-board.component.html',
  styleUrls: ['./game-board.component.scss']
})
export class GameBoardComponent implements OnInit {
  readonly boardMass = 8;
  board: BoardMass[];
  cover: BoardCover[];
  stone: Stone[];

  constructor() {
    this.board = [];
    this.cover = [];
    this.stone = [];
    for (let i = 0; i < this.boardMass * this.boardMass; i++) {
      this.board.push(this.makeMass(i));
      this.cover.push(this.makeCover(i));
    }
  }

  private makeMass(index: number): BoardMass {
    const x = (index % this.boardMass) * SIZE;
    const y = Math.floor(index / this.boardMass) * SIZE;
    const boardMass: BoardMass = { index, x, y, ...DEFAULT_BOARD };
    return boardMass;
  }

  private makeCover(index: number): BoardCover {
    const x = (index % this.boardMass) * SIZE;
    const y = Math.floor(index / this.boardMass) * SIZE;
    const boardCover: BoardCover = { index, x, y, ...DEFAULT_BOARD, opacity: 0 };
    boardCover.bgColor = "yellow";
    return boardCover;
  }

  private makeStone(index: number, isBlack: boolean): Stone {
    const x = (index % this.boardMass) * SIZE + SIZE / 2;
    const y = Math.floor(index / this.boardMass) * SIZE + SIZE / 2;
    const r = SIZE * 0.5 * 0.8;
    const stone: Stone = { index, x, y, r, stoneColor: isBlack ? "black" : "white", strokeColor: isBlack ? "white" : "black", strokeWidth: 3 };

    return stone;
  }

  ngOnInit(): void {
  }

  change(index: number) {
    if (this.stone.find(st => st.index == index)) {
      return;
    }
    this.changeCover(index, 0);
    this.stone.push(this.makeStone(index, true));
  }

  mouseover(index: number) {
    if (this.stone.find(st => st.index == index)) {
      return;
    }
    this.changeCover(index, 0.5);
  }

  mouseout(index: number) {
    this.changeCover(index, 0);
  }

  private changeCover(index: number, opacity: number) {
    const cover = this.cover.find(c => c.index == index);
    if (!cover) {
      return;
    }

    cover.opacity = opacity;
  }
}

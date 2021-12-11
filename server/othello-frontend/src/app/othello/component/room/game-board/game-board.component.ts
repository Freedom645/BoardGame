import { AfterViewInit, Component, ElementRef, EventEmitter, Input, OnInit, Output, ViewChild } from '@angular/core';
import { Point, StoneType } from 'src/app/othello/model/game';
import { ResizeService } from 'src/app/service/resize.service';

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
  strokeWidth: number,
}

@Component({
  selector: 'app-game-board',
  templateUrl: './game-board.component.html',
  styleUrls: ['./game-board.component.scss']
})
export class GameBoardComponent implements OnInit, AfterViewInit {
  @ViewChild('svg') svgElement!: ElementRef<SVGElement>;

  readonly boardMass = 8;
  boardPanel: BoardMass[];
  coverPanel: BoardCover[];
  stonePanel: Stone[];

  @Input() stone: StoneType[];
  @Output() clickBoard = new EventEmitter<Point>();

  size: number = 50;
  offset: number = 0;

  get DEFAULT_BOARD() {
    return {
      width: this.size - 2,
      height: this.size - 2,
      bgColor: "green"
    };
  }


  constructor(
    private resizeService: ResizeService,
  ) {
    this.boardPanel = [];
    this.coverPanel = [];
    this.stonePanel = [];
    this.stone = [];

    this.initializeBoard();
    this.resizeService.size.subscribe((_) => {
      this.resize();
    });
  }

  private initializeBoard() {
    for (let i = 0; i < this.boardMass * this.boardMass; i++) {
      this.boardPanel.push(this.makeMass(i));
      this.coverPanel.push(this.makeCover(i));
      this.stonePanel.push(this.makeStone(i));
    }
  }

  private resize() {
    const maxWidth = this.svgElement.nativeElement.clientWidth;
    this.size = Math.floor(maxWidth / this.boardMass);
    this.offset = Math.floor((maxWidth - this.size * this.boardMass) / 2);

    this.boardPanel.splice(0);
    this.coverPanel.splice(0);
    this.stonePanel.splice(0);
    for (let i = 0; i < this.boardMass * this.boardMass; i++) {
      this.boardPanel.push(this.makeMass(i));
      this.coverPanel.push(this.makeCover(i));
      this.stonePanel.push(this.makeStone(i));
    }
  }

  private makeMass(index: number): BoardMass {
    const x = (index % this.boardMass) * this.size + this.offset;
    const y = Math.floor(index / this.boardMass) * this.size + this.offset;
    const boardMass: BoardMass = { index, x, y, ...this.DEFAULT_BOARD };
    return boardMass;
  }

  private makeCover(index: number): BoardCover {
    const x = (index % this.boardMass) * this.size + this.offset;
    const y = Math.floor(index / this.boardMass) * this.size + this.offset;
    const boardCover: BoardCover = { index, x, y, ...this.DEFAULT_BOARD, opacity: 0 };
    boardCover.bgColor = "yellow";
    return boardCover;
  }

  private makeStone(index: number): Stone {
    const x = (index % this.boardMass) * this.size + this.size / 2 + this.offset;
    const y = Math.floor(index / this.boardMass) * this.size + this.size / 2 + this.offset;
    const r = this.size * 0.5 * 0.8;
    const stone: Stone = { index, x, y, r, strokeWidth: 3 };

    return stone;
  }

  ngOnInit(): void {
  }

  ngAfterViewInit(): void {
    this.resize();
  }

  change(index: number) {
    if (this.stone[index] != "None") {
      return;
    }

    const x = index % this.boardMass;
    const y = Math.floor(index / this.boardMass);
    this.clickBoard.emit({ x, y });

    this.stone[index] = Math.random() > 0.5 ? "White" : "Black";
    this.changeCover(index, 0);
  }

  mouseover(index: number) {
    if (this.stone[index] != "None") {
      return;
    }
    this.changeCover(index, 0.5);
  }

  mouseout(index: number) {
    this.changeCover(index, 0);
  }

  private changeCover(index: number, opacity: number) {
    const cover = this.coverPanel.find(c => c.index == index);
    if (!cover) {
      return;
    }

    cover.opacity = opacity;
  }
}

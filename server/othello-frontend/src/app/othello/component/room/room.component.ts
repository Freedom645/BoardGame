import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { WebSocketSubject } from 'rxjs/webSocket';
import { WebSocketService } from 'src/app/service/web-socket.service';
import { deserializer, GameMessage, GameStep, serializer, Step } from '../../model/message';
import { Board, Mass, Point, Stone, StoneType } from '../../model/game';
import { AccountService } from 'src/app/service/account.service';
import { Subscription } from 'rxjs';
import { RequestService } from '../../service/request.service';
import { Overlay } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { MatSpinner } from '@angular/material/progress-spinner';
import { GameLogicService } from '../../service/game-logic.service';
import { MatchingDialogComponent, MatchingDialogData } from './matching-dialog/matching-dialog.component';
import { MatDialog, MatDialogRef } from '@angular/material/dialog';

@Component({
  selector: 'app-room',
  templateUrl: './room.component.html',
  styleUrls: ['./room.component.scss']
})
export class RoomComponent implements OnInit, OnDestroy {
  roomId: string;
  playerName: string;
  board: Board;
  step: GameStep;

  private readonly subscription: Subscription[] = [];

  private overlayRef;

  constructor(
    private router: Router,
    private acRoute: ActivatedRoute,
    private accService: AccountService,
    private webSocketService: WebSocketService,
    private requestService: RequestService,
    private overlay: Overlay,
    private logic: GameLogicService,
    public dialog: MatDialog,
  ) {
    this.playerName = "";
    this.step = Step.Matching;
    this.roomId = acRoute.snapshot.params["id"];

    this.board = this.logic.newBoard();

    this.overlayRef = this.overlay.create({
      hasBackdrop: true,
      positionStrategy: this.overlay.position().global().centerHorizontally().centerVertically()
    });
  }

  ngOnInit(): void {
    // ユーザ名取得
    this.subscription.push(this.accService.getUsername().subscribe(name => this.playerName = name));
    // wsコネクト
    this.overlayRef.attach(new ComponentPortal(MatSpinner));
    this.webSocketService.connect<GameMessage>(this.roomId, serializer, deserializer)
      .subscribe(
        ws => {
          this.subscription.push(ws.subscribe(msg => this.receiveMessage(msg)));
          this.sendMessage = (msg) => ws.next(msg);
        }
      );
  }

  ngOnDestroy(): void {
    this.subscription.forEach(sub => sub.unsubscribe());
  }

  clickBoard(p: Point) {
    const value = this.requestService.requestGame(this.playerName, p);
    this.sendMessage(value);
  }

  /** メッセージ送信 */
  private sendMessage!: (msg: GameMessage) => void;

  /** メッセージ受信 */
  private receiveMessage(msg: GameMessage) {
    console.dir(msg);
    if (!msg.response) {
      return;
    }
    this.step = msg.response.step;
    this.board = msg.response.board;

    this.overlayRef.detach();
    switch (msg.response?.step) {
      case Step.Matching:
        this.matching(this.roomId);
        return;
      case Step.Pending:
        return;
      case Step.Waiting:
        return;
      case Step.Black:
        return;
      case Step.White:
        return;
      case Step.GameOver:
        return;
      case Step.Continue:
        return;
    }
  }

  private matchingDialog: MatDialogRef<MatchingDialogComponent, boolean> | null = null;
  private matching(id: string) {
    if (this.matchingDialog) {
      return;
    }
    this.matchingDialog = this.dialog.open(MatchingDialogComponent, {
      width: '100%',
      minHeight: 'calc(100vh - 90px)',
      height: 'auto',
      data: { roomId: id } as MatchingDialogData,
    });

    this.matchingDialog.afterClosed().subscribe(result => {
      if (result) {
        return;
      }
      this.router.navigate(["othello", "room"]);
      return;
    });
  }

  test() {
    const arr = [Step.Matching, Step.Pending, Step.Waiting, Step.Black, Step.White, Step.GameOver, Step.Continue];
    const index = arr.findIndex(st => st == this.step);
    this.step = arr[(index + 1) % arr.length];

    this.matching(this.roomId);
  }
}

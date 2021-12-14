import { Component, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { WebSocketSubject } from 'rxjs/webSocket';
import { WebSocketService } from 'src/app/service/web-socket.service';
import { deserializer, GameMessage, GameStep, PlayerType, serializer, Step } from '../../model/message';
import { Board, Mass, Point, Stone, StoneType } from '../../model/game';
import { AccountService } from 'src/app/service/account.service';
import { Subscription } from 'rxjs';
import { RequestService } from '../../service/request.service';
import { Overlay } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { MatSpinner } from '@angular/material/progress-spinner';
import { GameLogicService } from '../../service/game-logic.service';
import { MatStepper } from '@angular/material/stepper';

@Component({
  selector: 'app-room',
  templateUrl: './room.component.html',
  styleUrls: ['./room.component.scss']
})
export class RoomComponent implements OnInit, OnDestroy {
  readonly roomId: string;
  readonly playerType: PlayerType;
  playerName: string;
  opponentPlayer: string;

  board: Board;
  step: GameStep;

  @ViewChild('stepper') private stepper!: MatStepper;
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
  ) {
    this.playerName = "";
    this.opponentPlayer = "";

    this.step = Step.Pending;
    this.roomId = acRoute.snapshot.params["id"];
    this.playerType = acRoute.snapshot.queryParams["pt"] != "spectator" ? "player" : "spectator";

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
    this.webSocketService.connect<GameMessage>(`/room/${this.roomId}/ws?pt=${this.playerType}`, serializer, deserializer)
      .subscribe(
        ws => {
          this.subscription.push(ws.subscribe(msg => this.receiveMessage(msg)));
          this.sendMessage = (msg) => ws.next(msg);
        }
      );
  }

  ngOnDestroy(): void {
    this.subscription.forEach(sub => sub.unsubscribe());
    this.overlayRef.detach();
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
      case Step.Pending:
        return;
      case Step.Black:
        return;
      case Step.White:
        return;
      case Step.GameOver:
        return;
    }
  }

  test() {
    const arr = [Step.Pending, Step.Black, Step.White, Step.GameOver];
    const index = arr.findIndex(st => st == this.step);
    this.step = arr[(index + 1) % arr.length];

    this.stepper.next();
  }
}

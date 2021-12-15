import { Component, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { WebSocketSubject } from 'rxjs/webSocket';
import { WebSocketService } from 'src/app/service/web-socket.service';
import { deserializer, GameMessage, GameStep, PlayerType, serializer, Step } from '../../model/message';
import { Board, Mass, Point, Stone, StoneType } from '../../model/game';
import { AccountService } from 'src/app/service/account.service';
import { Subscription } from 'rxjs';
import { RequestService } from '../../service/request.service';
import { Overlay, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { MatSpinner } from '@angular/material/progress-spinner';
import { GameLogicService } from '../../service/game-logic.service';
import { MatStepper } from '@angular/material/stepper';
import { Player, Turn } from '../../model/room';

@Component({
  selector: 'app-room',
  templateUrl: './room.component.html',
  styleUrls: ['./room.component.scss']
})
export class RoomComponent implements OnInit, OnDestroy {
  readonly roomId: string;
  readonly playerType: PlayerType;
  playerName: string;

  owner!: Player;
  opponent!: Player;
  private turn!: Turn;

  board: Board;
  step: GameStep;
  stepIndex: number;
  blackName: string;
  whiteName: string;

  @ViewChild('stepper') private stepper!: MatStepper;
  private readonly subscription: Subscription[] = [];

  private overlayRef: OverlayRef;

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

    this.step = Step.Pending;
    this.stepIndex = 0;
    this.roomId = acRoute.snapshot.params["id"];
    this.playerType = acRoute.snapshot.queryParams["pt"] != "spectator" ? "player" : "spectator";

    this.blackName = "";
    this.whiteName = "";

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

  clickApprove(isApprove: boolean) {
    const value = this.requestService.requestPending(isApprove);
    this.sendMessage(value);
  }

  clickBoard(p: Point) {
    const value = this.requestService.requestGame(p);
    this.sendMessage(value);
  }

  /** メッセージ送信 */
  private sendMessage!: (msg: GameMessage) => void;

  /** メッセージ受信 */
  private receiveMessage(msg: GameMessage) {
    if (!msg.response) {
      return;
    }
    this.step = msg.response.step;
    this.board = msg.response.board;
    this.owner = msg.response.owner;
    this.turn = msg.response.turn;
    this.setPlayers(msg.response.players);

    if (this.step == "pending") {
      this.stepIndex = 0;
    } else if (this.step == "black" || this.step == "white") {
      this.stepIndex = 1;
    } else if (this.step == "gameOver") {
      this.stepIndex = 2;
    }

    this.overlayRef.detach();
  }

  private setPlayers(players: Player[]) {
    if (!players) {
      return;
    }
    this.opponent = players[1];
    this.blackName = players.find(p => p.id == this.turn.blackId)?.name ?? "";
    this.whiteName = players.find(p => p.id == this.turn.whiteId)?.name ?? "";
  }

}

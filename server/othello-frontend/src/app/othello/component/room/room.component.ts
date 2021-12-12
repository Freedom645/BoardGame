import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { WebSocketSubject } from 'rxjs/webSocket';
import { WebSocketService } from 'src/app/service/web-socket.service';
import { deserializer, GameMessage, serializer } from '../../model/message';
import { Point, StoneType } from '../../model/game';
import { AccountService } from 'src/app/service/account.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-room',
  templateUrl: './room.component.html',
  styleUrls: ['./room.component.scss']
})
export class RoomComponent implements OnInit, OnDestroy {
  roomId: string;
  playerName: string;
  stone: StoneType[];

  private wsSubject!: WebSocketSubject<GameMessage>;
  private readonly subscription: Subscription[] = [];

  constructor(
    private router: ActivatedRoute,
    private accService: AccountService,
    private webSocketService: WebSocketService
  ) {
    this.playerName = "";
    this.roomId = router.snapshot.params["id"];
    this.stone = new Array<StoneType>(8 * 8).fill("none");

    this.webSocketService.connect<GameMessage>(this.roomId, serializer, deserializer).subscribe(
      ws => {
        this.wsSubject = ws;
        this.subscription.push(this.wsSubject.subscribe(msg => this.receiveMessage(msg)));
      }
    );
  }

  ngOnInit(): void {
    this.subscription.push(this.accService.getUsername().subscribe(name => this.playerName = name));
  }

  ngOnDestroy(): void {
    this.subscription.forEach(sub => sub.unsubscribe());
    this.wsSubject.unsubscribe();
  }

  clickBoard(point: Point) {
    const value: GameMessage = {
      request: {
        playerName: this.playerName,
        game: {
          point: point
        }
      }
    };
    this.wsSubject.next(value);
  }

  private receiveMessage(msg: GameMessage) {
    console.dir(msg);
  }
}

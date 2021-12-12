import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { map, Subject } from 'rxjs';
import { WebSocketSubject } from 'rxjs/webSocket';
import { WebSocketService } from 'src/app/service/web-socket.service';
import { GameRequestMessage, GameResponseMessage, Point, StoneType } from '../../model/game';

@Component({
  selector: 'app-room',
  templateUrl: './room.component.html',
  styleUrls: ['./room.component.scss']
})
export class RoomComponent implements OnInit {
  roomId: string;
  stone: StoneType[];

  private wsSubject!: WebSocketSubject<string>;

  constructor(
    private router: ActivatedRoute,
    private webSocketService: WebSocketService
  ) {
    this.roomId = router.snapshot.params["id"];
    this.stone = new Array(8 * 8).fill("None");

    this.webSocketService.connect<GameRequestMessage>(this.roomId).subscribe(
      ws => {
        this.wsSubject = ws;
        this.wsSubject.subscribe(msg => {
          console.dir(msg);
        });
      }
    );
  }

  ngOnInit(): void {

  }

  clickBoard(point: Point) {
    this.wsSubject.next(`[${point.x},${point.y}]`);
  }

}

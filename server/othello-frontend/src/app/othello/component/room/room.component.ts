import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { map, Subject } from 'rxjs';
import { WebSocketService } from 'src/app/service/web-socket.service';
import { GameRequestMessage, GameResponseMessage } from '../../model/game';

@Component({
  selector: 'app-room',
  templateUrl: './room.component.html',
  styleUrls: ['./room.component.scss']
})
export class RoomComponent implements OnInit {
  roomId: string;

  private wsSubject;

  constructor(
    private router: ActivatedRoute,
    private webSocketService: WebSocketService
  ) {
    this.roomId = router.snapshot.params["id"];

    this.wsSubject = this.webSocketService.connect<GameRequestMessage>(this.roomId);
    this.wsSubject.subscribe((res) => {
      console.log(res);
    });

  }

  ngOnInit(): void {

  }

  onClick() {
    this.wsSubject.next("test");
  }

}

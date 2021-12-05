import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RoomListComponent } from './component/room-list/room-list.component';
import { ApiService } from '../service/api.service';
import { WebSocketService } from '../service/web-socket.service';
import { MaterialModule } from '../material.module';
import { RoomDetailComponent } from './component/room-detail/room-detail.component';
import { OthelloRoutingModule } from './othello-routing.module';
import { RoomComponent } from './component/room/room.component';
import { GameBoardComponent } from './component/room/game-board/game-board.component';



@NgModule({
  declarations: [
    RoomListComponent,
    RoomDetailComponent,
    RoomComponent,
    GameBoardComponent
  ],
  imports: [
    CommonModule,
    MaterialModule,
    OthelloRoutingModule,
  ],
  providers: [
    ApiService,
    WebSocketService,
  ],
})
export class OthelloModule { }

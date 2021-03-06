import { NgModule } from '@angular/core';
import { CommonModule, DatePipe } from '@angular/common';
import { RoomListComponent } from './component/room-list/room-list.component';
import { ApiService } from '../service/api.service';
import { WebSocketService } from '../service/web-socket.service';
import { MaterialModule } from '../material.module';
import { RoomDetailComponent } from './component/room-detail/room-detail.component';
import { OthelloRoutingModule } from './othello-routing.module';
import { RoomComponent } from './component/room/room.component';
import { GameBoardComponent } from './component/room/game-board/game-board.component';
import { RoomTableComponent } from './component/room-list/room-table/room-table.component';
import { ResizeService } from '../service/resize.service';
import { RequestService } from './service/request.service';
import { GameLogicService } from './service/game-logic.service';
import { GameStepComponent } from './component/room/game-step/game-step.component';
import { MatchingComponent } from './component/room/matching/matching.component';
import { GameOverComponent } from './component/room/game-over/game-over.component';



@NgModule({
  declarations: [
    RoomListComponent,
    RoomDetailComponent,
    RoomComponent,
    GameBoardComponent,
    RoomTableComponent,
    GameStepComponent,
    MatchingComponent,
    GameOverComponent,
  ],
  imports: [
    CommonModule,
    MaterialModule,
    OthelloRoutingModule,
  ],
  providers: [
    ApiService,
    WebSocketService,
    ResizeService,
    RequestService,
    GameLogicService,
    DatePipe,
  ],
})
export class OthelloModule { }

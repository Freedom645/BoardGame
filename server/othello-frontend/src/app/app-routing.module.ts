import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { GameComponent } from './component/game/game.component';
import { RoomComponent } from './component/room/room.component';

const routes: Routes = [
  {path: "room", component: RoomComponent },
  {path: "game", component: GameComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

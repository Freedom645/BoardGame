import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RoomListComponent } from './component/room-list/room-list.component';
import { RoomComponent } from './component/room/room.component';

const routes: Routes = [
  { path: "room", component: RoomListComponent },
  { path: "room/:id", component: RoomComponent },
  { path: '', redirectTo: 'room', pathMatch: 'full'}
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class OthelloRoutingModule { }

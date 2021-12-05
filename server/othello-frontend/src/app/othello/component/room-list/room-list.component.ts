import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ActivatedRoute, Router } from '@angular/router';
import { ApiService } from 'src/app/service/api.service';
import { Room } from '../../model/room';
import { RoomDetailComponent, RoomDetailDialogDate } from '../room-detail/room-detail.component';

@Component({
  selector: 'app-room-list',
  templateUrl: './room-list.component.html',
  styleUrls: ['./room-list.component.scss']
})
export class RoomListComponent implements OnInit {

  displayedColumns: string[] = ['roomId', 'created', 'button'];
  roomList: Room[] = []

  constructor(
    private api : ApiService,
    private router: Router,
    private acRoute: ActivatedRoute,
    public dialog: MatDialog ,
  ) { }

  ngOnInit(): void {
    this.api.getRoomList().subscribe(
      (res) => this.roomList = res
    )
  }

  onClickRoomJoin(id: string){
    const dialogRef = this.dialog.open(RoomDetailComponent, {
      width: '100%',
      minHeight: 'calc(100vh - 90px)',
      height : 'auto',
      data: {roomId: id} as RoomDetailDialogDate,
    });

    dialogRef.afterClosed().subscribe(result => {
      console.dir(result)
      if(result){
        this.router.navigate([id], {relativeTo: this.acRoute})
      }
    });
  }
}

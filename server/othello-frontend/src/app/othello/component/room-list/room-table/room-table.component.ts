import { Component, Input, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ActivatedRoute, Router } from '@angular/router';
import { Room } from 'src/app/othello/model/room';
import { ApiService } from 'src/app/service/api.service';
import { RoomDetailComponent, RoomDetailDialogDate } from '../../room-detail/room-detail.component';

@Component({
  selector: 'app-room-table',
  templateUrl: './room-table.component.html',
  styleUrls: ['./room-table.component.scss']
})
export class RoomTableComponent implements OnInit {

  displayedColumns: string[] = ['roomId', 'created', 'button'];
  @Input() roomList: Room[] = [];

  constructor(
    private router: Router,
    private acRoute: ActivatedRoute,
    public dialog: MatDialog,
  ) { }

  ngOnInit(): void {
  }

  onClickRoomJoin(id: string) {
    const dialogRef = this.dialog.open(RoomDetailComponent, {
      width: '100%',
      minHeight: 'calc(100vh - 90px)',
      height: 'auto',
      data: { roomId: id } as RoomDetailDialogDate,
    });

    dialogRef.afterClosed().subscribe(result => {
      console.dir(result);
      if (result) {
        this.router.navigate([id], { relativeTo: this.acRoute });
      }
    });
  }
}

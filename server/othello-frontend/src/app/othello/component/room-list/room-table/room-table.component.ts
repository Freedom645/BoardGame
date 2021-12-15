import { Component, Input, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ActivatedRoute, Router } from '@angular/router';
import { Room } from 'src/app/othello/model/room';
import { ApiService } from 'src/app/service/api.service';
import { RoomDetailComponent, RoomDetailDialogData, RoomDetailDialogOut } from '../../room-detail/room-detail.component';

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
    const dialogRef = this.dialog.open<RoomDetailComponent, RoomDetailDialogData, RoomDetailDialogOut>(RoomDetailComponent, {
      width: '100%',
      minHeight: 'calc(30vh)',
      height: 'auto',
      data: { roomId: id } as RoomDetailDialogData,
    });

    dialogRef.afterClosed().subscribe(result => {
      if (!result) {
        return;
      }
      switch (result.buttonType) {
        case "join":
          this.router.navigate([id], { relativeTo: this.acRoute, queryParams: { pt: "player" } });
          return;
        case "spectate":
          this.router.navigate([id], { relativeTo: this.acRoute, queryParams: { pt: "spectator" } });
          return;
      }
    });
  }
}

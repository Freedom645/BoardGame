import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { ApiService } from 'src/app/service/api.service';
import { Room } from '../../model/room';

export interface RoomDetailDialogDate {
  roomId: string
}

@Component({
  selector: 'app-room-detail',
  templateUrl: './room-detail.component.html',
  styleUrls: ['./room-detail.component.scss']
})
export class RoomDetailComponent implements OnInit {
  room: Room = {id: this.data.roomId, created: new Date()}

  constructor(
    private api: ApiService,
    private dialogRef: MatDialogRef<RoomDetailComponent>,
    @Inject(MAT_DIALOG_DATA) private data: RoomDetailDialogDate,
  ) {}


  ngOnInit(): void {
    this.api.getRoom(this.data.roomId).subscribe(res => {
      this.room = res
    })
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

}

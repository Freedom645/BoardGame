import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { Room } from 'src/app/othello/model/room';
import { ApiService } from 'src/app/service/api.service';

export interface MatchingDialogData {
  roomId: string;

}

@Component({
  selector: 'app-matching-dialog',
  templateUrl: './matching-dialog.component.html',
  styleUrls: ['./matching-dialog.component.scss']
})
export class MatchingDialogComponent implements OnInit {

  room: Room = { id: this.data.roomId, created: new Date() };

  constructor(
    private api: ApiService,
    private dialogRef: MatDialogRef<MatchingDialogComponent>,
    @Inject(MAT_DIALOG_DATA) private data: MatchingDialogData,
  ) { }


  ngOnInit(): void {
    this.api.getRoom(this.data.roomId).subscribe(res => {
      this.room = res;
    });
  }

  approve(): void {

  }

}

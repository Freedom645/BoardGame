import { DatePipe } from '@angular/common';
import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { Subscription } from 'rxjs';
import { ApiService } from 'src/app/service/api.service';


export interface RoomDetailDialogData {
  roomId: string;
}

export type ButtonType = "return" | "join" | "spectate";
export interface RoomDetailDialogOut {
  buttonType: ButtonType;
}

interface Item {
  label: string;
  value: string;
}

@Component({
  selector: 'app-room-detail',
  templateUrl: './room-detail.component.html',
  styleUrls: ['./room-detail.component.scss']
})
export class RoomDetailComponent implements OnInit {

  private readonly subscriptions: Subscription[] = [];

  items: Item[] = [];

  constructor(
    private api: ApiService,
    private dialogRef: MatDialogRef<RoomDetailComponent, RoomDetailDialogOut>,
    private datePipe: DatePipe,
    @Inject(MAT_DIALOG_DATA) private data: RoomDetailDialogData,
  ) {
    this.items = [
      { label: "部屋ID", value: "" },
      { label: "部屋主", value: "" },
      { label: "作成日時", value: "" },
    ];
  }


  ngOnInit(): void {
    this.subscriptions.push(
      this.api.getRoom(this.data.roomId).subscribe(res => {
        this.items[0].value = res.id.substring(0, 8);
        this.items[1].value = res.owner.name ?? "unknown";
        this.items[2].value = this.datePipe.transform(res.created, "yyyy/MM/dd HH:mm:ss") ?? "";
      }));
  }

  clickButton(buttonType: ButtonType): void {
    this.dialogRef.close({ buttonType: buttonType });
  }

}

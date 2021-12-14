import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Player } from 'src/app/othello/model/room';

@Component({
  selector: 'app-matching',
  templateUrl: './matching.component.html',
  styleUrls: ['./matching.component.scss']
})
export class MatchingComponent implements OnInit {

  @Input() roomId: string = "";
  @Input() owner!: Player;
  @Input() opponent!: Player;

  @Output() clickApprove = new EventEmitter<boolean>();

  constructor() {

  }

  ngOnInit(): void {
  }

  approve(isApprove: boolean) {
    if (this.opponent) {
      this.clickApprove.emit(isApprove);
    }
  }

}

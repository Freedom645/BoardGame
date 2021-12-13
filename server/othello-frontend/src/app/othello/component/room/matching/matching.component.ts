import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';

@Component({
  selector: 'app-matching',
  templateUrl: './matching.component.html',
  styleUrls: ['./matching.component.scss']
})
export class MatchingComponent implements OnInit {

  @Input() roomId: string = "";
  @Input() playerName: string = "";
  @Input() opponentPlayer: string = "";

  @Output() clickApprove = new EventEmitter<boolean>();

  constructor() {

  }

  ngOnInit(): void {
  }

  approve(isApprove: boolean) {
    if (this.opponentPlayer) {
      this.clickApprove.emit(isApprove);
    }
  }

}

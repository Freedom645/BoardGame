import { Component, Input, OnChanges, OnInit, SimpleChanges } from '@angular/core';
import { GameStep, Step } from 'src/app/othello/model/message';

@Component({
  selector: 'app-game-step',
  templateUrl: './game-step.component.html',
  styleUrls: ['./game-step.component.scss']
})
export class GameStepComponent implements OnInit, OnChanges {

  @Input() step: GameStep;
  displayStep: string;

  constructor() {
    this.step = Step.Pending;
    this.displayStep = "";
  }

  ngOnInit(): void {
  }

  ngOnChanges(changes: SimpleChanges): void {
    if (changes["step"]) {
      this.displayStep = this.convertDisplayStep(this.step);
    }
  }

  private convertDisplayStep(step: GameStep): string {
    switch (step) {
      case Step.Pending: return "Pending";
      case Step.Black: return "Black Turn";
      case Step.White: return "White Turn";
      case Step.GameOver: return "Game Over";
    }
    return "";
  }
}

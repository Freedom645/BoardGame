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
    this.step = Step.Matching;
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
      case Step.Matching: return "Matching";
      case Step.Pending: return "Pending";
      case Step.Black: return "Black Turn";
      case Step.White: return "White Turn";
      case Step.Waiting: return "Waiting";
      case Step.GameOver: return "Game Over";
      case Step.Continue: return "Continue";
    }
    return "";
  }
}

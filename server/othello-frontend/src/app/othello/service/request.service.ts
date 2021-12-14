import { Injectable } from '@angular/core';
import { Point } from '../model/game';
import { GameMessage } from '../model/message';

@Injectable()
export class RequestService {

  constructor() { }

  public requestGame(point: Point): GameMessage {
    return {
      request: {
        game: {
          point: point
        }
      }
    };
  }

  public requestPending(isApprove: boolean): GameMessage {
    return {
      request: {
        pending: {
          isApproved: isApprove
        }
      }
    };
  }

  public requestGameOver(isContinue: boolean): GameMessage {
    return {
      request: {
        gameOver: {
          isContinued: isContinue
        }
      }
    };
  }
}

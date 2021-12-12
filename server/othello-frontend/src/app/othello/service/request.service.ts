import { Injectable } from '@angular/core';
import { Point } from '../model/game';
import { GameMessage } from '../model/message';

@Injectable()
export class RequestService {

  constructor() { }

  public requestGame(playerName: string, point: Point): GameMessage {
    const value: GameMessage = {
      request: {
        playerName: playerName,
        game: {
          point: point
        }
      }
    };
    return value;
  }
}

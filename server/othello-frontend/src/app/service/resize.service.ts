import { Injectable } from '@angular/core';
import { EventManager } from '@angular/platform-browser';
import { map, Observable, Subject } from 'rxjs';

export interface Size {
  width: number;
  height: number;
}

@Injectable()
export class ResizeService {
  private eventSubject = new Subject<UIEvent>();

  constructor(private eventManager: EventManager) {
    // Windowに対するペーストイベントをSubjectにする
    this.eventManager.addGlobalEventListener('window', 'resize', async (event: UIEvent) => this.eventSubject.next(event));
  }

  get event(): Observable<UIEvent> {
    return this.eventSubject;
  }

  get size(): Observable<Size> {
    return this.eventSubject.pipe(map(() => {
      return { 'width': window.innerWidth, 'height': window.innerHeight };
    }));
  }
}

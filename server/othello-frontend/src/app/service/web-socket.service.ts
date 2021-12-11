import { Subject, Observable, Observer } from 'rxjs';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { webSocket, WebSocketSubject, WebSocketSubjectConfig } from "rxjs/webSocket";

@Injectable()
export class WebSocketService {

  private readonly URL = environment.api.ws + environment.api.host;

  public connect<T>(roomId: string): WebSocketSubject<T> {
    return webSocket<T>(this.makeConfig<T>(`${this.URL}/room/${roomId}/ws`));
  }

  private makeConfig<T>(url: string): WebSocketSubjectConfig<T> {
    const config: WebSocketSubjectConfig<T> = {
      url: url
    };
    return config;
  }

}

import { Subject, Observable, Observer, concatMap, map } from 'rxjs';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { webSocket, WebSocketSubject, WebSocketSubjectConfig } from "rxjs/webSocket";
import { AccountService } from './account.service';

export type Serializer<T> = (value: T) => string;
export type Deserializer<T> = (e: MessageEvent) => T;

@Injectable()
export class WebSocketService {

  private readonly URL = environment.api.ws + environment.api.host;

  constructor(
    private acc: AccountService,
  ) { }

  public connect<T>(roomId: string, serializer: Serializer<T>, deserializer: Deserializer<T>): Observable<WebSocketSubject<T>> {
    return this.acc.getJwt()
      .pipe(
        map(jwt => this.makeConfig<T>(`${this.URL}/room/${roomId}/ws`, jwt, serializer, deserializer)),
        map(config => webSocket<T>(config))
      );
  }

  private makeConfig<T>(url: string, jwt: string, serializer: Serializer<T>, deserializer: Deserializer<T>): WebSocketSubjectConfig<T> {
    const config: WebSocketSubjectConfig<T> = {
      url: url,
      serializer: serializer,
      deserializer: deserializer,
      protocol: jwt
    };
    return config;
  }


}

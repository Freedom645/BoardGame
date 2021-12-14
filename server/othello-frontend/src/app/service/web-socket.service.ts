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

  public connect<T>(path: string, serializer: Serializer<T>, deserializer: Deserializer<T>, ...protocol: string[]): Observable<WebSocketSubject<T>> {
    return this.acc.getJwt()
      .pipe(
        map(jwt => this.makeConfig<T>(`${this.URL}${path}`, serializer, deserializer, [jwt, ...protocol])),
        map(config => webSocket<T>(config))
      );
  }

  private makeConfig<T>(url: string, serializer: Serializer<T>, deserializer: Deserializer<T>, protocol: string[]): WebSocketSubjectConfig<T> {
    const config: WebSocketSubjectConfig<T> = {
      url: url,
      serializer: serializer,
      deserializer: deserializer,
      protocol: protocol,
    };
    return config;
  }


}

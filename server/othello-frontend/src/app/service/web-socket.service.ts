import { Subject, Observable, Observer, concatMap, map } from 'rxjs';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { webSocket, WebSocketSubject, WebSocketSubjectConfig } from "rxjs/webSocket";
import { AccountService } from './account.service';

@Injectable()
export class WebSocketService {

  private readonly URL = environment.api.ws + environment.api.host;

  constructor(
    private acc: AccountService,
  ) { }

  public connect<T>(roomId: string): Observable<WebSocketSubject<T>> {
    return this.acc.getJwt()
      .pipe(
        map(jwt => this.makeConfig<T>(`${this.URL}/room/${roomId}/ws`, jwt)),
        map(config => webSocket<T>(config))
      );
  }

  private makeConfig<T>(url: string, jwt: string): WebSocketSubjectConfig<T> {
    const config: WebSocketSubjectConfig<T> = {
      url: url,
      serializer: this.serializer,
      deserializer: this.deserializer,
      protocol: jwt
    };
    return config;
  }

  private serializer = (value: any) => {
    console.log("serializer");
    console.dir(value);
    return JSON.stringify(value);
  };
  private deserializer = (e: MessageEvent) => {
    console.log("deserializer");
    console.dir(e.data);
    return JSON.parse(e.data || "null");
  };

}

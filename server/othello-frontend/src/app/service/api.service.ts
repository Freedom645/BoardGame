import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { concatMap, map, Observable } from 'rxjs';
import { Room } from '../othello/model/room';
import { environment } from 'src/environments/environment';
import { AccountService } from './account.service';

@Injectable()
export class ApiService {

  private readonly URL = environment.api.http + environment.api.host;

  constructor(
    private http: HttpClient,
    private acc: AccountService,
  ) { }

  private makeHeader(): Observable<HttpHeaders> {
    return this.acc.getJwt().pipe(map(jwt => {
      const header = new HttpHeaders()
        .append('Authorization', `Bearer ${jwt}`);
      return header;
    }));
  }

  private get<T>(url: string): Observable<T> {
    return this.makeHeader().pipe(
      concatMap(headers => this.http.get<T>(url, { headers: headers }))
    );
  }

  private post<T>(url: string, body: any): Observable<T> {
    return this.makeHeader().pipe(
      concatMap(headers => this.http.post<T>(url, body, { headers: headers }))
    );
  }

  /** 部屋一覧を取得 */
  public getRoomList(): Observable<Room[]> {
    return this.get<Room[]>(`${this.URL}/room`);
  }

  /** 部屋情報を取得 */
  public getRoom(id: string): Observable<Room> {
    return this.get<Room>(`${this.URL}/room/${id}`);
  }

  /** 部屋を作成 */
  public createRoom(body?: any): Observable<Room> {
    return this.post<Room>(`${this.URL}/room`, body ?? {});
  }
}

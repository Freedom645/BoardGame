import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Room } from '../othello/model/room';
import { environment } from 'src/environments/environment';

@Injectable()
export class ApiService {

  private readonly URL = environment.api.http + environment.api.host
  private readonly HEADERS = this.makeHeader()

  constructor(
    private http: HttpClient
  ) { }

  private makeHeader(): HttpHeaders {
    const header = new HttpHeaders()
    header.append('Content-Type', 'application/x-www-form-urlencoded')
    return new HttpHeaders()
  }

  private get<T>(url: string) {
    return this.http.get<T>(url, {headers: this.HEADERS})
  }

  /** 部屋一覧を取得 */
  public getRoomList(): Observable<Room[]> {
    return this.get<Room[]>(`${this.URL}/room`)
  }

  /** 部屋情報を取得 */
  public getRoom(id: string): Observable<Room>{
    return this.get<Room>(`${this.URL}/room/${id}`)
  }
}

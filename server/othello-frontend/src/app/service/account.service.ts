import { Injectable } from '@angular/core';
import { AngularFireAuth } from '@angular/fire/compat/auth';
import { concatMap, map, from, Observable, take, of } from 'rxjs';

@Injectable()
export class AccountService {

  constructor(
    private afAuth: AngularFireAuth
  ) { }

  public isLoggedIn(): Observable<boolean> {
    return this.afAuth.authState.pipe(
      take(1),
      map((user) => {
        if (user) {
          return true;
        }
        return false;
      }));
  }

  public getUsername(): Observable<string> {
    return this.afAuth.user.pipe(
      map((user) => user?.displayName),
      concatMap((name) => name ? of(name) : this.getUid().pipe(map(uid => "Guest_" + uid.substring(0, 8))))
    );
  }

  public getUid(): Observable<string> {
    return this.afAuth.user.pipe(
      map((user) => user?.uid ?? "")
    );
  }

  public logout(): Promise<void> {
    return this.afAuth.signOut();
  }

  public getJwt(): Observable<string> {
    return this.afAuth.authState.pipe(concatMap((user) => {
      return from(user?.getIdToken() ?? new Promise<string>(re => ""));
    }
    ));
  }
}



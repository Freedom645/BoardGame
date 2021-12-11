import { Injectable } from '@angular/core';
import { AngularFireAuth } from '@angular/fire/compat/auth';
import { map, Observable, take } from 'rxjs';

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
      map((user) => user?.displayName ?? "Guest")
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
}

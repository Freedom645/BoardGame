import { Injectable } from '@angular/core';
import { AngularFireAuth } from '@angular/fire/compat/auth';
import { CanActivate, Router } from '@angular/router';
import { map, Observable, take } from 'rxjs';
import { AccountService } from './account.service';

@Injectable()
export class AuthGuardService implements CanActivate {

  constructor(
    private accService: AccountService,
    private router: Router,
  ) { }

  canActivate(): Observable<boolean> {
    return this.accService.isLoggedIn().pipe(map((res) => {
      if (!res) {
        this.router.navigate(["login"]);
      }
      return res;
    }));
  }

}

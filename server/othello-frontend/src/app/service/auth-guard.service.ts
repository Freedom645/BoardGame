import { Injectable } from '@angular/core';
import { AngularFireAuth } from '@angular/fire/compat/auth';
import { ActivatedRouteSnapshot, CanActivate, Router, UrlTree } from '@angular/router';
import { map, Observable, take } from 'rxjs';
import { AccountService } from './account.service';

@Injectable()
export class AuthGuardService implements CanActivate {

  private loginUrlTree: UrlTree;
  private redirectUrlTree: UrlTree;

  constructor(
    private accService: AccountService,
    private router: Router,
  ) {
    this.loginUrlTree = this.router.parseUrl("login");
    this.redirectUrlTree = this.router.parseUrl("othello");
  }

  canActivate(route: ActivatedRouteSnapshot): Observable<UrlTree | boolean> {
    return this.accService.isLoggedIn()
      .pipe(map((isLoggedIn) => {
        // by https://stackoverflow.com/questions/48197067/how-can-you-use-angulars-canactivate-to-negate-the-result-of-a-guard
        if (isLoggedIn) {
          if (route.routeConfig?.path == "login") {
            return this.redirectUrlTree;
          }
          return true;
        }
        if (route.routeConfig?.path == "login") {
          return true;
        }
        return this.loginUrlTree;
      }));
  }

}

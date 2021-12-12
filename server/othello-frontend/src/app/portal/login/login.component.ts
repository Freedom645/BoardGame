import { ThrowStmt } from '@angular/compiler';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FirebaseUISignInSuccessWithAuthResult, FirebaseUISignInFailure } from 'firebaseui-angular';
import { filter } from 'rxjs';
import { AccountService } from 'src/app/service/account.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  constructor(
    private accService: AccountService,
    private route: Router,
  ) { }

  ngOnInit(): void {
    this.accService.getUid().pipe(filter((user) => !!user)).subscribe((user) => this.route.navigate(["othello", "room"]));
  }

  successCallback(signInSuccessData: FirebaseUISignInSuccessWithAuthResult) {
    this.route.navigate(["othello", "room"]);
  }

  async errorCallback(errorData: FirebaseUISignInFailure) {
    // ログインエラー
    console.log(errorData);
  }

}

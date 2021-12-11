import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { Router } from '@angular/router';
import { AccountService } from 'src/app/service/account.service';

@Component({
  selector: 'app-site-header',
  templateUrl: './site-header.component.html',
  styleUrls: ['./site-header.component.scss']
})
export class SiteHeaderComponent implements OnInit {

  @Output() clickSideMenu = new EventEmitter<any>();

  constructor(
    private accService: AccountService,
    private router: Router
  ) { }

  ngOnInit(): void {
  }

  navigate(path: string) {
    this.router.navigateByUrl(path);
  }

  logout() {
    this.accService.logout().then(() => this.router.navigate(["login"]));
  }

}

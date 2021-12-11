import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-top',
  templateUrl: './top.component.html',
  styleUrls: ['./top.component.scss']
})
export class TopComponent implements OnInit {

  constructor(
    private router: Router,
    private acRoute: ActivatedRoute,
  ) { }

  ngOnInit(): void {
  }

  onClickStart() {
    this.router.navigate(["login"]);
  }
}

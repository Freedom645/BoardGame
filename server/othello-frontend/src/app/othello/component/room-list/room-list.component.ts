import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription, take } from 'rxjs';
import { AccountService } from 'src/app/service/account.service';
import { ApiService } from 'src/app/service/api.service';
import { Room } from '../../model/room';

@Component({
  selector: 'app-room-list',
  templateUrl: './room-list.component.html',
  styleUrls: ['./room-list.component.scss']
})
export class RoomListComponent implements OnInit, OnDestroy {

  readonly subscriptions: Subscription[] = [];

  userName: string = "";
  uid: string = "";
  roomList: Room[] = [];

  constructor(
    private accService: AccountService,
    private apiService: ApiService,
    private router: Router,
    private acRouter: ActivatedRoute,
  ) {
  }

  ngOnInit(): void {
    this.subscriptions.push(this.accService.getUsername().subscribe(userName => this.userName = userName));
    this.subscriptions.push(this.accService.getUid().subscribe(uid => this.uid = uid));
    this.subscriptions.push(this.apiService.getRoomList().pipe(take(1)).subscribe((res) => this.roomList = res));
  }

  ngOnDestroy(): void {
    this.subscriptions.forEach(sub => sub.unsubscribe());
  }

  createRoom() {
    this.subscriptions.push(this.apiService.createRoom().pipe(take(1)).subscribe((room) => {
      this.router.navigate([room.id], { relativeTo: this.acRouter });
    }));
  }

  test() {
    this.accService.getJwt();
  }

}

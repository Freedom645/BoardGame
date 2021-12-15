import { Overlay, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { Component, OnDestroy, OnInit } from '@angular/core';
import { MatSpinner } from '@angular/material/progress-spinner';
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

  private readonly subscriptions: Subscription[] = [];

  private readonly overlayRef: OverlayRef;

  userName: string = "";
  uid: string = "";
  roomList: Room[] = [];

  constructor(
    private accService: AccountService,
    private apiService: ApiService,
    private router: Router,
    private acRouter: ActivatedRoute,
    private overlay: Overlay,
  ) {
    this.overlayRef = this.overlay.create({
      hasBackdrop: true,
      positionStrategy: this.overlay.position().global().centerHorizontally().centerVertically()
    });
  }

  ngOnInit(): void {
    this.refreshRoom();
    this.subscriptions.push(this.accService.getUsername().subscribe(userName => this.userName = userName));
    this.subscriptions.push(this.accService.getUid().subscribe(uid => this.uid = uid));
  }

  ngOnDestroy(): void {
    this.subscriptions.forEach(sub => sub.unsubscribe());
    this.overlayRef.detach();
  }

  /** 部屋を作成して移動する */
  createRoom() {
    this.subscriptions.push(this.apiService.createRoom().pipe(take(1)).subscribe((room) => {
      this.router.navigate([room.id], { relativeTo: this.acRouter });
    }));
  }

  /** 部屋一覧を取得 */
  refreshRoom() {
    this.overlayRef.attach(new ComponentPortal(MatSpinner));
    this.roomList = [];
    this.subscriptions.push(this.apiService.getRoomList().pipe(take(1)).subscribe((res) => {
      this.roomList = res;
      this.overlayRef.detach();
    }));
  }

}

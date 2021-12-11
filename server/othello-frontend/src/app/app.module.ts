import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { OthelloModule } from './othello/othello.module';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { SiteFrameComponent } from './component/site-frame/site-frame.component';
import { MaterialModule } from './material.module';
import { SiteFooterComponent } from './component/site-footer/site-footer.component';
import { SiteHeaderComponent } from './component/site-header/site-header.component';
import { PortalModule } from './portal/portal.module';
import { AuthGuardService } from './service/auth-guard.service';
import { AccountService } from './service/account.service';

@NgModule({
  declarations: [
    AppComponent,
    SiteFrameComponent,
    SiteFooterComponent,
    SiteHeaderComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    BrowserAnimationsModule,
    MaterialModule,
    OthelloModule,
    PortalModule,
  ],
  providers: [
    AuthGuardService,
    AccountService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }

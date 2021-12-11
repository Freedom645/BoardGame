import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuardService } from './service/auth-guard.service';

const routes: Routes = [
  { path: 'othello', loadChildren: () => import('./othello/othello-routing.module').then(mod => mod.OthelloRoutingModule), canActivate: [AuthGuardService] },
  { path: '', loadChildren: () => import('./portal/portal-routing.module').then(mod => mod.PortalRoutingModule) },
  { path: '**', redirectTo: 'othello', pathMatch: 'full' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

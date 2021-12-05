import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  { path: 'othello', loadChildren: () => import('./othello/othello-routing.module').then(mod => mod.OthelloRoutingModule)},
  { path: '**', redirectTo: 'othello', pathMatch: 'full'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

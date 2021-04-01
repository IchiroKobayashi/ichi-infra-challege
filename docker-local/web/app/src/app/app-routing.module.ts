import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  { path:`scraping-challenge`, loadChildren: () => import(`./features/scraping-challenge/scraping-challenge.module`).then(m => m.ScrapingChallengeModule) }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

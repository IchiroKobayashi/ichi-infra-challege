import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { TitleScrapingComponent } from './title-scraping/page/title-scraping.component';

const routes: Routes = [
  { path:`title-scraping`, component: TitleScrapingComponent, data: {} }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ScrapingChallengeRoutingModule { }

import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from "@angular/common";
import { RouterModule } from "@angular/router";

import { ScrapingChallengeRoutingModule } from './scraping-challenge-routing.module';
import { TitleScrapingComponent } from './title-scraping/page/title-scraping.component';
import { TitleScrapingService } from './title-scraping/service/title-scraping.service';


@NgModule({
  declarations: [
    // Product Components
    TitleScrapingComponent
  ],
  imports: [
    // Angular Common Modules
    CommonModule,
    RouterModule,
    ReactiveFormsModule,
    FormsModule,
    // Product Routing Modules
    ScrapingChallengeRoutingModule
  ],
  entryComponents: [
    // Components Generated Dynamically
  ],
  providers: [
    // Product Services (DI)
    TitleScrapingService
  ],
  bootstrap: [],
  schemas: [CUSTOM_ELEMENTS_SCHEMA]
})
export class ScrapingChallengeModule { }

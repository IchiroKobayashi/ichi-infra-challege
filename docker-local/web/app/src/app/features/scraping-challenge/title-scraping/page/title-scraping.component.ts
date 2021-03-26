import { Component, OnInit, OnDestroy, ViewEncapsulation } from '@angular/core';
import { ActivatedRoute, Router } from "@angular/router";
import { BehaviorSubject, Subject, Subscription } from "rxjs";
import { FormBuilder, FormsModule, FormControl } from '@angular/forms';
import { TitleScrapingService } from '../service/title-scraping.service';
import { TitleScrapingEntity } from '../model/title-scraping.model';


@Component({
  selector: 'app-title-scraping',
  templateUrl: './title-scraping.component.html',
  styleUrls: ['./title-scraping.component.less'],
  encapsulation: ViewEncapsulation.None
})
export class TitleScrapingComponent implements OnInit, OnDestroy {

  // Constructor DI
  constructor(
    private service: TitleScrapingService,
    private route: ActivatedRoute,// To receive parameters
    private router: Router,// For Transition
    private formBuilder: FormBuilder,
  ) { }

  // Property Definition
  destroyed$ = new Subject();
  private subscriptions: Array<Subscription> = [];
  titles: Array<TitleScrapingEntity>;
  isTitle: boolean = false;
  urls: string;

  ngOnDestroy(): void {
    this.destroyed$.next();
    this.destroyed$.complete();
    this.subscriptions.forEach(removeSubscription => removeSubscription.unsubscribe());
  }

  ngOnInit(): void {
    this.titles = [];
  }

  getTitles(): void {
    this.titles = [];
    this.subscriptions.push(
      this.service.getTitles(this.urls).subscribe(response => {
        if(response.length > 0) {
          this.isTitle = true;
        }
        this.titles = response;
      })
    );
  }

}

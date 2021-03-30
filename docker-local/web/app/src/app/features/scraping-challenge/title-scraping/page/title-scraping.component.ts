import { Component, OnInit, OnDestroy, ViewEncapsulation } from '@angular/core';
import { ActivatedRoute, Router } from "@angular/router";
import { BehaviorSubject, Observable, Subject, Subscription } from "rxjs";
import { FormBuilder, FormsModule, FormControl } from '@angular/forms';
import { TitleScrapingService } from '../service/title-scraping.service';
import { TitleScrapingEntity } from '../model/title-scraping.model';
import { TEXT } from '../../../../../resources/texts/features/scraping-challenge/title-scraping/text';

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
  texts: { [key:string]: string};
  pageReady: boolean = false;

  ngOnDestroy(): void {
    this.destroyed$.next();
    this.destroyed$.complete();
    this.subscriptions.forEach(removeSubscription => removeSubscription.unsubscribe());
  }

  ngOnInit(): void {
    this.subscriptions.push(
      TEXT().subscribe(res =>{
        this.texts = res;
      })
    );
    this.titles = [];
    this.pageReady = true;
  }

  getTitles(): void {
    this.titles = [];
    this.isTitle = false;
    this.subscriptions.push(
      this.service.getTitles(this.urls).subscribe(response => {
        if(response.length > 0) {
          this.isTitle = true;
        }
        this.titles = response;
      })
    );
  }

  notifyMe(): void {
    this.service.notifyMe().subscribe(response => {
      console.log(response);
    })

  }

}

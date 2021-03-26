import {ComponentFixture, ComponentFixtureAutoDetect, TestBed} from '@angular/core/testing';
import { FormsModule } from '@angular/forms';
import { TitleScrapingComponent } from './title-scraping.component';
import { TitleScrapingService } from '../service/title-scraping.service';

describe('TitleScrapingComponent', () => {
  let component: TitleScrapingComponent;
  let fixture: ComponentFixture<TitleScrapingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TitleScrapingComponent ],
      imports:[ FormsModule ],
      providers:[ TitleScrapingService,
        { provide: ComponentFixtureAutoDetect, useValue: true }]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TitleScrapingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

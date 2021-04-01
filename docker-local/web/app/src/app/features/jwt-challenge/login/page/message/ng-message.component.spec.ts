import { ComponentFixture, ComponentFixtureAutoDetect, TestBed } from '@angular/core/testing';
import { FormsModule } from '@angular/forms';
import { NgMessageComponent } from './ng-message.component';

describe('NgMessageComponent', () => {
  let component: NgMessageComponent;
  let fixture: ComponentFixture<NgMessageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NgMessageComponent ],
      imports:[ FormsModule ],
      providers:[
        { provide: ComponentFixtureAutoDetect, useValue: true }]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NgMessageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

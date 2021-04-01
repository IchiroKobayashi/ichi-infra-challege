import {Component, OnInit, OnDestroy, ViewEncapsulation, NgModule, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from "@angular/router";
import { BehaviorSubject, Observable, Subject, Subscription } from "rxjs";
import { FormGroup, FormBuilder, FormsModule, FormControl, Validators, AbstractControl, ValidatorFn } from '@angular/forms';
import { LoginService } from '../service/login.service';
import { UserLoginEntity } from '../model/user-login.model';
import { NgMessageComponent } from './message/ng-message.component';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.less'],
  encapsulation: ViewEncapsulation.None
})
export class LoginComponent implements OnInit, OnDestroy {
  @ViewChild('NgMessageComponent')
  private ngMessageComponent:NgMessageComponent;

  // Constructor DI
  constructor(
    private service: LoginService,
    private route: ActivatedRoute,// To receive parameters
    private router: Router,// For Transition
    private formBuilder: FormBuilder,
  ) { }


  // Property Definition
  destroyed$ = new Subject();
  private subscriptions: Array<Subscription> = [];
  pageReady: boolean = true;
  public formGroup: FormGroup;

  ngOnDestroy(): void {
    this.destroyed$.next();
    this.destroyed$.complete();
    this.subscriptions.forEach(removeSubscription => removeSubscription.unsubscribe());
  }

  ngOnInit(): void {
    // this.pageReady = true;
    this.formGroup = new FormGroup({
      email: new FormControl('', [
        Validators.email, // e-mailフォーマットチェック
        Validators.required
        // this.duplicateEmailValidator() // 任意のバリデーション(今回は重複チェック)
      ]),
      password: new FormControl('', [
        Validators.required,
        Validators.minLength(8) // 最低8文字
      ])
    });
    console.log(this.formGroup);
  }

  public get email(): FormControl {
    return this.formGroup.get('email') as FormControl;
  }

  public get password(): FormControl {
    return this.formGroup.get('password') as FormControl;
  }

  onSubmit(event: Event) {
    // TODO: Use EventEmitter with form value
    console.log(event);
    // this.service.login(event);
  }

  // private duplicateEmailValidator(): ValidatorFn {
  //   return (control: AbstractControl): {[key: string]: any} | null => {
  //     const ngList = [
  //       'hoge@example.com',
  //       'huga@test.co.jp'
  //     ];
  //     const duplicate = ngList.includes(control.value);
  //     return duplicate ? {'duplicateEmail': {value: control.value}} : null;
  //   };
  // }
}

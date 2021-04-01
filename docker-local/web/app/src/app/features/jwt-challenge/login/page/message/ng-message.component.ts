import { Component, OnInit, Input } from '@angular/core';
import { FormGroup, FormControl, AbstractControl } from '@angular/forms';

@Component({
  selector: 'app-ng-message',
  templateUrl: './ng-message.component.html',
  styleUrls: ['./ng-message.component.less']
})
export class NgMessageComponent implements OnInit {
  @Input() formGroup: FormGroup;
  email: FormControl;
  password: FormControl;

  constructor() { }

  ngOnInit() {
    this.email = this.formGroup.controls.email as FormControl;
    this.password = this.formGroup.controls.password as FormControl;
  }
}

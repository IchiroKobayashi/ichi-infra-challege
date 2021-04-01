import { Component, OnInit, Input } from '@angular/core';
import { FormControl } from '@angular/forms';

@Component({
  selector: 'app-ng-message',
  templateUrl: './ng-message.component.html',
  styleUrls: ['./ng-message.component.less']
})
export class NgMessageComponent implements OnInit {
  @Input() form: FormControl;

  constructor() { }

  ngOnInit() {
  }
}

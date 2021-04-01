import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from "@angular/common";
import { RouterModule } from "@angular/router";

import { JwtChallengeRoutingModule } from './jwt-challenge-routing.module';
import { LoginComponent } from './login/page/login.component';
import { LoginService } from './login/service/login.service';
import { NgMessageComponent } from './login/page/message/ng-message.component';

@NgModule({
  declarations: [
    // Product Components
    LoginComponent,
    NgMessageComponent
  ],
  imports: [
    // Angular Common Modules
    CommonModule,
    RouterModule,
    ReactiveFormsModule,
    FormsModule,
    // Product Routing Modules
    JwtChallengeRoutingModule
  ],
  entryComponents: [
    // Components Generated Dynamically
  ],
  providers: [
    // Product Services (DI)
    LoginService
  ],
  bootstrap: [],
  schemas: [CUSTOM_ELEMENTS_SCHEMA]
})
export class JwtChallengeModule { }

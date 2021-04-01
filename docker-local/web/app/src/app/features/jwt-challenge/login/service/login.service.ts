import { Observable } from "rxjs";
import { Injectable} from "@angular/core";
import { HttpClient} from "@angular/common/http";
import { UserLoginEntity } from '../model/user-login.model';
import { JwtChallengeConst } from '../../constant/jwt-challenge-const';

@Injectable()
export class LoginService {

  constructor(
    private http: HttpClient
  ) {
  }

  login(email: string, password: string): Observable<UserLoginEntity>{
    return this.http.post<UserLoginEntity>(
      JwtChallengeConst.API_URL + 'login',
      {
        email: email,
        password: password
      }
    )
  }
}

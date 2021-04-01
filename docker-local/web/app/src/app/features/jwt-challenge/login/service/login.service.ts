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

  login(urls: string): Observable<Array<UserLoginEntity>>{
    return this.http.post<Array<UserLoginEntity>>(
      JwtChallengeConst.API_URL + 'login',
      {
        params:{
          urls: urls
        }
      }
    )
  }
}

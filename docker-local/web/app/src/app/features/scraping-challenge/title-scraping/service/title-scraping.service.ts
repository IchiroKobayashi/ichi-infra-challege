import { Observable } from "rxjs";
import { Injectable} from "@angular/core";
import { HttpClient} from "@angular/common/http";
import { TitleScrapingEntity } from '../model/title-scraping.model';
import { ScrapingChallengeConst } from '../../constant/scraping-challenge-const';

@Injectable()
export class TitleScrapingService {

  constructor(
    private http: HttpClient
  ) {
  }

  getTitles(urls: string): Observable<Array<TitleScrapingEntity>>{
    return this.http.get<Array<TitleScrapingEntity>>(
      ScrapingChallengeConst.API_URL + 'scrape',
      {
        params:{
          urls: urls
        }
      }
    )
  }

}

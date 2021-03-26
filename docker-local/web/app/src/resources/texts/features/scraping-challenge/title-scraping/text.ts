import { Observable, of} from "rxjs";

export const TEXT = (): Observable<{ [key: string]: string}> => {
  return of<{ [key:string]: string}>(
      {
        'title': 'Webサイトからtitleをスクレイピングするアプリ',
        'formDesc' :'URL：[,]カンマ区切りで複数入力'
      }
  )
}

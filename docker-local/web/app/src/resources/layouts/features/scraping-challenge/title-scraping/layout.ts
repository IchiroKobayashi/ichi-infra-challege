import { Observable, of} from "rxjs";

export const LAYOUT = (): Observable<Array<object>> => {
  return of<Array<object>>(
    [
      {
        'title': 'Test Application',
        'formDesc' :'URL：[,]カンマ区切りで複数入力'
      }
    ]
  )
}

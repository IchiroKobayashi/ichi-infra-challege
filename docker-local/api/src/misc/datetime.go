package misc

import (
	"time"

	"github.com/go-openapi/strfmt"
)

const (
	timeformat     = "2006-01-02 15:04:05"
	dateformat     = "2006/01/02"
	yyyymmddformat = "20060102"
)

// Now 現在時刻を UTC で返します
func Now() time.Time {
	return time.Now().UTC()
}

// MaxDate 最大日付を UTC で返します(仮で2100/12/31 14:59:59)
func MaxDate() time.Time {
	return time.Date(2100, time.December, 31, 14, 59, 59, 0, time.UTC)
}

// NowJST 現在時刻を JST で返します
func NowJST() time.Time {
	nowUTC := Now()
	return nowUTC.In(JSTLocation())
}

// JSTLocation 日本時間のタイムゾーンを返す
func JSTLocation() *time.Location {
	return time.FixedZone("Asia/Tokyo", 9*60*60)
}

// Between  target が、from と to の間にあるかを返します。（同じ時刻も含む）
func Between(target, from, to time.Time) bool {
	return (target.Equal(from) || target.After(from)) && (target.Equal(to) || target.Before(to))
}

// NowJSTYear 現在日付（日本時間）の年を取得(yyyy)
func NowJSTYear() string {

	day := NowJST()
	const layout = "2006"
	return day.Format(layout)
}

// MonthLastDate 該当日付の末日を返す
func MonthLastDate(t time.Time) time.Time {
	//翌月の1日から1日引く方法で実装

	// UTC時間で 15時以降だったら翌日として計算
	if t.Hour() >= 15 {
		t = t.AddDate(0, 0, 1)
	}

	return time.Date(t.Year(), t.Month()+1, 1, 14, 59, 59, 0, time.UTC).AddDate(0, 0, -1)

}

// NowToUnixMilliSecound 現在時刻をミリ秒
func NowToUnixMilliSecound() int64 {
	return time.Now().UTC().UnixNano()
}

// ToSwaggerDateTime strfmt.DateTime 型に変換します
func ToSwaggerDateTime(candidate *time.Time) strfmt.DateTime {
	if candidate != nil {
		dt := strfmt.DateTime(*candidate)
		return dt
	}
	return strfmt.NewDateTime()
}

// ToTime strfmt.DateTime型をTime型に変換します。
func ToDateTimeSwagger(d *strfmt.DateTime) *time.Time {
	if d != nil {
		t, e := time.Parse(time.RFC3339, d.String())
		if e == nil {
			return &t
		}
	}
	return nil
}

// ToSwaggerDate strfmt.Date 型に変換します
func ToSwaggerDate(candidate time.Time) strfmt.Date {
	dt := strfmt.Date(candidate)
	return dt
}

// ToTime strfmt.Date型をTime型に変換します。
func ToDateSwagger(d *strfmt.Date) *time.Time {
	if d != nil {
		t, e := time.Parse(strfmt.RFC3339FullDate, d.String())
		if e == nil {
			return &t
		}
	}
	return nil
}

// ToTime strfmt.Date型をTime型に変換します。
func ToDateValueSwagger(d strfmt.Date) *time.Time {
	if t, e := time.Parse(strfmt.RFC3339FullDate, d.String()); e == nil {
		return &t
	}
	return nil
}

// ParseDateTime 文字列を日付に変換して返します
func ParseDateTime(candidate string) *time.Time {
	dt, err := time.Parse(timeformat, candidate)
	if err != nil {
		return nil
	}
	return &dt
}

// ParseString (ParseDateTimeと同じフォーマットで文字列を返す)
func ParseString(t time.Time) string {
	return t.Format(timeformat)
}

// StringToDBTime "2006/01/02" の日本時間を DB 保存用の UTC 時間に変換して返します
func StringToDBTime(strDate string) *time.Time {
	j, err := time.ParseInLocation(dateformat, strDate, JSTLocation())
	if err != nil {
		return nil
	}
	u := j.In(time.UTC)
	return &u
}

// JstParseString 日付を日本時間の文字列に変換して返します
func JstParseString(t time.Time) string {
	timeJST := t.In(JSTLocation())
	return timeJST.Format(timeformat)
}

func unixMilli(t time.Time) int64 {
	return t.Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

// MakeTimestampMilli (Unix時間にミリ秒を加えた13桁の整数)
func MakeTimestampMilli() int64 {
	return unixMilli(Now())
}

// TimeToJstStringPt1 日付を日本時間の文字列に変換して返します
// func TimeToJstStringPt1(t time.Time) string {

// 	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

// 	timeJST := t.In(jst)

// 	return timeJST.Format(timeformat2)
// }

// ToJST JSTにする
func ToJST(t time.Time) time.Time {
	// すでに JST だったらそのまま
	if t.Location().String() == JSTLocation().String() {
		return t
	}
	// JST にして返す
	timeInJST := t.In(JSTLocation())
	return timeInJST
}

// ToUTC UTCにする
func ToUTC(t time.Time) time.Time {
	// すでに UTC だったらそのまま
	if t.Location().String() == time.UTC.String() {
		return t
	}
	// UTC にして返す
	timeInUTC := t.In(time.UTC)
	return timeInUTC

}

// ToDateText yyyy/mm/dd の文字列にする
func ToDateText(t time.Time) string {
	return t.Format(dateformat)
}

// Toyyyymmdd yyyymmdd の文字列にする
func Toyyyymmdd(t time.Time) string {
	return t.Format(yyyymmddformat)
}

// FirstDayOfMonth 指定した日が入っている月の1日の0時を取得。タイムゾーンは変えない。
func FirstDayOfMonth(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), 1 /*firstDay*/, 0, 0, 0, 0, d.Location())
}

// MaxDayOfMonth 指定した日が入っている月の最終日の最終時刻（23:59:59）を取得。タイムゾーンは変えない。
func MaxDayOfMonth(t time.Time) time.Time {
	// 指定した日の入っている月の次の月の1日の1秒前を取得
	return time.Date(t.Year(), t.Month()+1 /**/, 1 /*firstDay*/, 0, 0, 0, 0, t.Location()).Add(-1 * time.Second)
}

// DateTimeToPtrDateTime strfmt.UUIDのアドレスを返す
func DateTimeToPtrDateTime(d strfmt.DateTime) *strfmt.DateTime {
	return &d
}

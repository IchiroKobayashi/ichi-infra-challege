package misc

import (
	"crypto/rand"
	"encoding/binary"
	"reflect"
	"regexp"
	"strconv"

	"github.com/go-openapi/swag"
)

// RandSeed rund に渡すシードを生成します
func RandSeed() int64 {
	var seed int64
	err := binary.Read(rand.Reader, binary.LittleEndian, &seed)
	if err != nil {
		panic(err)
	}
	return seed
}

// ToInt64 インターフェイス型を int64 に
func ToInt64(value interface{}) int64 {
	r := reflect.ValueOf(value)
	if r.IsValid() {
		switch r.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return r.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return int64(r.Uint())
		case reflect.Float32, reflect.Float64:
			return int64(r.Float())
		}
	}
	return 0
}

// ToUint8 インターフェイス型を uint8 に
func ToUint8(value interface{}) uint8 {
	r := reflect.ValueOf(value)
	if r.IsValid() {
		switch r.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return uint8(r.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return uint8(r.Uint())
		case reflect.Float32, reflect.Float64:
			return uint8(r.Float())
		}
	}
	return 0
}

var commas = regexp.MustCompile(`(\d+)(\d{3})`)

// FormatCommas 3桁おきにカンマを挿入します
func FormatCommas(num int) string {
	str := strconv.Itoa(num)
	for i := 0; i < (len(str)-1)/3; i++ {
		str = commas.ReplaceAllString(str, "$1,$2")
	}
	return str
}

//I32 *int32をintに変換
func I32(i *int32) int {
	return int(swag.Int32Value(i))
}

//I64 *int64をintに変換
func I64(i *int64) int {
	return int(swag.Int64Value(i))
}

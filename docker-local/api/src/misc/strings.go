package misc

import (
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-openapi/strfmt"
)

func init() {
	rand.Seed(RandSeed())
}

// RandomString returns randomized string
func RandomString(letters []byte, n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// InsertChar group ごとに char を挿入した文字列を返します
func InsertChar(value, char string, group int) string {
	re := regexp.MustCompile(fmt.Sprintf("([^%s]+)([^%s]{%d})", char, char, group))
	for {
		hyphend := re.ReplaceAllString(value, fmt.Sprintf("$1%s$2", char))
		if hyphend == value {
			return hyphend
		}
		value = hyphend
	}
}

// ToString インターフェイス型を文字列に
func ToString(value interface{}) string {
	if candidate, ok := value.(string); ok {
		return candidate
	}
	return ""
}

var nonDigits = regexp.MustCompile("[^0-9]")

// DigitsOnly 数字のみを残した文字列を返します
func DigitsOnly(value string) string {
	return nonDigits.ReplaceAllString(value, "")
}

// TypeOf インターフェイスの型を文字列で返します
func TypeOf(value interface{}) string {
	ref := reflect.ValueOf(value)
	if ref.IsValid() {
		switch kind := ref.Kind(); kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return "Integer"
		case reflect.Float32, reflect.Float64:
			return "Float"
		case reflect.Complex64, reflect.Complex128:
			return "Complex"
		case reflect.Ptr, reflect.Uintptr, reflect.UnsafePointer:
			return "Pointer"
		case reflect.Bool, reflect.Array, reflect.Slice,
			reflect.Map, reflect.Chan, reflect.Func,
			reflect.Interface, reflect.String, reflect.Struct:
			return kind.String()
		}
	}
	return "Invalid"
}

// Contains スライスに指定の文字列を含むか
func Contains(records []string, item string) bool {
	for _, record := range records {
		if strings.EqualFold(item, record) {
			return true
		}
	}
	return false
}

// CutUTF8StringByCount (主にDB用) UTF-8 の文字列を、指定の文字数を超えない長さに調整する。
func CutUTF8StringByCount(s string, size int) string {
	count := len([]rune(s))
	if count <= size {
		return s
	}
	return string([]rune(s)[0:size])
}

// UUIDToPtrUUID strfmt.UUIDのアドレスを返す
func UUIDToPtrUUID(s strfmt.UUID) *strfmt.UUID {
	return &s
}

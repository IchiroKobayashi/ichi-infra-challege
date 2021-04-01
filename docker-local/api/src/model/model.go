package model

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"time"
	"crypto/sha256"

	"github.com/ichi-infra-challenge/docker-local/api/src/config"
	"github.com/jinzhu/gorm"
)

// User ユーザー情報のテーブル情報
type User struct {
	ID        string       `json:"id" gorm:"column:id;"`
	Email     string       `json:"email" gorm:"column:email;"`
	Password  string       `json:"password" gorm:"column:password;"`
	Username  string       `json:"username" gorm:"column:username;"`
	CreatedAt time.Time    `json:"createdAt" gorm:"column:created_at;"`
	UpdatedAt time.Time    `json:"updatedAt" gorm:"column:updated_at;"`
}

// Create ユーザー情報のテーブル情報
func Create(name string) (sql.Result, error) {
	// db, err := sql.Open("mysql", config.MySQLConnection)
	db, err := sql.Open("mysql", "local:local@tcp(db:3306)/infra-challenge?charset=utf8&parseTime=true")

	fmt.Println(err)
	fmt.Println(db)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
	fmt.Println(name)

	sqlParams := []interface{}{name}
	query := `INSERT INTO users (username) VALUES (?)`

	fmt.Println(query)
	defer db.Close()
	res, err := db.Exec(query, sqlParams...)
	if err != nil {
		return res, err
	}
	return res, err

}

func getSHA256Binary(s string) []byte {
	r := sha256.Sum256([]byte(s))
	return r[:]
}


// FindUser だよ
func FindUserByEmailPass(email string, password string) (User, error) {
	// db, err := sql.Open("mysql", "local:local@tcp(127.0.0.1:3306)/infra-challenge")
	db, err := sql.Open("mysql", config.MySQLConnection)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
	passSha256b := getSHA256Binary(password)
	passSha256Str := hex.EncodeToString(passSha256b)
	fmt.Println(passSha256Str)

	sqlParams := []interface{}{ email, passSha256Str }
	query := "Select id, username FROM users WHERE email = ? and password = ?"

	fmt.Println(query)

	defer db.Close()
	var id string
	var username string
	err = db.QueryRow(query, sqlParams...).Scan(&id, &username)
	fmt.Println(id)
	fmt.Println(username)
	if err != nil {
		log.Fatal(err)
	}
	user := User{ID: id, Username: username}

	return user, err
}

// GetAll するよ
func GetAll() ([]User, error) {
	// db, err := sql.Open("mysql", config.MySQLConnection)
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	}
	var users []User
	db.Order("id").Find(&users)

	db.Close()

	return users, err
}

// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "local"
	PASS := "local"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "infra-challenge"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}

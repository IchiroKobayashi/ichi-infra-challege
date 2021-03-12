package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ichi-infra-challenge/docker-local/api/src/config"
	"github.com/jinzhu/gorm"
)

// User ユーザー情報のテーブル情報
type User struct {
	ID        int       `json:"id" gorm:"column:id;"`
	Name      string    `json:"name" gorm:"column:name;"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;"`
}

// Create ユーザー情報のテーブル情報
func Create(name string) (sql.Result, error) {
	// db, err := sql.Open("mysql", config.MySQLConnection)
	db, err := sql.Open("mysql", "local:local@tcp(127.0.0.1:3306)/infra-challenge?charset=utf8&parseTime=true")

	fmt.Println(err)
	fmt.Println(db)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
	fmt.Println(name)

	sqlParams := []interface{}{name}
	query := `INSERT INTO users (name) VALUES (?)`

	fmt.Println(query)
	defer db.Close()
	res, err := db.Exec(query, sqlParams...)
	if err != nil {
		return res, err
	}
	return res, err

}

// FindByName だよ
func FindByName(name string) (bool, error) {
	// db, err := sql.Open("mysql", "local:local@tcp(127.0.0.1:3306)/infra-challenge")
	db, err := sql.Open("mysql", config.MySQLConnection)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
	fmt.Println(name)

	sqlParams := []interface{}{name}
	query := "Select name FROM `users` WHERE `name` = ?"

	fmt.Println(query)

	defer db.Close()
	err = db.QueryRow(query, sqlParams...).Scan(&name)
	if err != nil {
		return false, err
	}
	fmt.Println(name)

	if name != "" {
		return true, err
	}

	return false, err
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
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "infra-challenge"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}

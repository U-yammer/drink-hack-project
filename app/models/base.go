package models

import (
    "crypto/sha1"
    "database/sql"
    "drink_hack_project/config"
    "fmt"
    "github.com/google/uuid"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

var Db *sql.DB

var err error

const (
    tableNameUser    = "users"
    tableNameTodo    = "todos"
    tableNameSession = "sessions"
    tableNameDrink   = "drinks"
    tableNameDInfo   = "dInfos"
)

func init() {
    Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
    if err != nil {
        log.Fatalln(err)
    }

    cmdU := fmt.Sprintf(` CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

    _, err := Db.Exec(cmdU)
    if err != nil {
        log.Fatalln(err)
    }

    cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`, tableNameTodo)
    Db.Exec(cmdT)

    cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		user_id INTEGER,
		created_at DATETIME)`, tableNameSession)
    Db.Exec(cmdS)

    _, err = Db.Exec(cmdT)
    if err != nil {
        log.Fatalln(err)
    }

    cmdD := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		drink_name STRING,
		amount INTEGER,
		created_at DATETIME)`, tableNameDrink)
    _, err = Db.Exec(cmdD)
    if err != nil {
        log.Fatalln(err)
    }

}

/* uuid(ユーザを一意に識別するID) を生成する関数 */
func createUUID() (uuidObject uuid.UUID) {
    uuidObject, _ = uuid.NewUUID()
    return uuidObject
}

/* ユーザの Password を生成する関数．ハッシュ値 (不可逆的な値をつくる) */
func Encrypt(plaintext string) (cryptext string) {
    cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
    return cryptext
}

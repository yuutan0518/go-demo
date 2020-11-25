package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yuutan0518/go-demo/entity"
)

const (
	// データベース
	Dialect = "mysql"

	// ユーザー名
	DBUser = "root"

	// パスワード
	DBPass = "password"

	// プロトコル*(コンテナのサーバにアクセスする場合は、コンテナ名を記載)
	DBProtocol = "tcp(mysql:3306)"

	// DB名
	DBName = "anger_log"
)

var (
	db  *gorm.DB
	err error
)

// Init is inirialize db from main function
func Init() {

	connectTemplate := "%s:%s@%s/%s"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)
	db, err := gorm.Open(Dialect, connect)

	if err != nil {
		log.Println(err.Error())
	}

	db.AutoMigrate(&entity.User{})

}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		log.Println(err.Error())
	}
}

// AutoMigration
// func autoMigration() {
// 	db.AutoMigrate(&entity.User{})
// }

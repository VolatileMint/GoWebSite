package model

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	// .envからdb情報を取得
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("BD_HOST")
	port := os.Getenv("DB_PORT")
	database_name := os.Getenv("DB_DATABASE_NAME")
	// [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
	dns := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error(), db)
	}
	//db.Migrator().DropTable(&User{}) // テーブルを削除する
	// テーブルがない場合は新規作成する(もしくはすでにある場合は消して作り直す)
	if !db.Migrator().HasTable(&User{}) {
		db.Migrator().DropTable(&User{})
		db.Migrator().CreateTable(&User{})
		fmt.Println("テーブルを新規で作成")
	}
	return db
}

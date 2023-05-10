package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

// データベース接続の初期化
func InitDatabase() error {
	// .envファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// データベースへの接続設定
	// 環境変数からデータベースの接続情報を取得
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// DSNの作成
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	// 既に初期化済みの場合は何もしない
	if db != nil {
		return nil
	}

	// データベース接続の確立
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	// 接続確認
	err = conn.Ping()
	if err != nil {
		return err
	}

	// グローバル変数にデータベース接続を設定
	db = conn

	log.Println("Database connected")

	return nil
}

func GetDB() *sql.DB {
	// データベース接続が初期化されていない場合は初期化する
	if db == nil {
		err := InitDatabase()
		if err != nil {
			log.Fatal(err)
		}
	}

	return db
}

package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFIle string) {
	logfile, err := os.OpenFile(logFIle, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) // ファイルの読み書き，作成，追記を指定
	if err != nil {
		log.Fatalln(err)
	}

	multiLogFile := io.MultiWriter(os.Stdout, logfile) // 標準出力 ＆ ログファイルに出力する
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // ログのフォーマットを指定
	log.SetOutput(multiLogFile) // セットアップ！
}
package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// ログファイル読み込み
	// 引数：ファイルのパス、フラグ（読み書き|新規ファイル作成|追記）、パーミッション(基本0666でOK)
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// ログファイル読み込み時問題が起きたら、エラー内容を表示して処理を終了する。
	if err != nil {
		log.Fatalf("file=logfile err=%s", err.Error())
	}
	// 引数で複数出力先を指定して、まとめて書き込みする。
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	// フラグ設定（日付|時刻|ソースファイル（ファイル名のみ））
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// ログを出力
	log.SetOutput(multiLogFile)
}

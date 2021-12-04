package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var clients = make(map[uuid.UUID]*websocket.Conn) // 接続されるクライアント
var broadcast = make(chan GameResponseMessage)    // メッセージブロードキャストチャネル

// アップグレーダ
var upgrader = websocket.Upgrader{}

func main() {
	// ファイルサーバーを立ち上げる
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	// websockerへのルーティングを紐づけ
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()
	// サーバーをlocalhostのポート8000で立ち上げる
	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	// エラーがあった場合ロギングする
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

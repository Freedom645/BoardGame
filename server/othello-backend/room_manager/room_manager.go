package room_manager

import (
	"log"
	"net/http"
	"sync"

	"github.com/Freedom645/BoardGame/room"
	"github.com/Freedom645/BoardGame/room/player"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type RoomManager struct {
	m      map[uuid.UUID]*room.Room
	locker sync.RWMutex
}

type GameRequestMessage struct {
	UUID     uuid.UUID `json:"uuid"`
	UserName string    `json:"userName"`
}

type GameResponseMessage struct {
}

/* 部屋一覧 */
var managerInstance = &RoomManager{m: make(map[uuid.UUID]*room.Room)}

/* メッセージブロードキャストチャネル */
var broadcast = make(chan GameResponseMessage)

/* アップグレーダ */
var upgrader = websocket.Upgrader{}

func CreateRoom() (*room.Room, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	newRoom := room.NewRoom(id)

	// 排他処理
	managerInstance.locker.Lock()
	defer managerInstance.locker.Unlock()

	managerInstance.m[newRoom.UUID()] = newRoom

	return newRoom, nil
}

func JoinFirstPlayer(uuid uuid.UUID, player player.Player) bool {
	// 書き込みロック
	managerInstance.locker.RLock()
	defer managerInstance.locker.Unlock()

	room := managerInstance.m[uuid]

	return room.SetFirstPlayer(player)
}

/* コネクション確立 */
func handleConnections(w http.ResponseWriter, r *http.Request) {
	// 送られてきたGETリクエストをwebsocketにアップグレード
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// 関数が終わった際に必ずwebsocketのコネクションを閉じる
	defer ws.Close()

	// クライアントを新しく登録
	clients[ws] = true

	for {
		var msg Message
		// 新しいメッセージをJSONとして読み込みMessageオブジェクトにマッピングする
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// 新しく受信されたメッセージをブロードキャストチャネルに送る
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// ブロードキャストチャネルから次のメッセージを受け取る
		msg := <-broadcast
		// 現在接続しているクライアント全てにメッセージを送信する
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

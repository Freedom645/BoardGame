package controller

import (
	"net/http"
	"sync"

	"github.com/Freedom645/BoardGame/controller/model/room_model"
	"github.com/Freedom645/BoardGame/domain/room"
	"github.com/Freedom645/BoardGame/service/room_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gopkg.in/olahol/melody.v1"
)

/* 部屋 */
type socketStruct struct {
	melody *melody.Melody
	room   *room.Room
	locker sync.RWMutex
}

/* WebSocketコネクション */
type melodyManagerStruct struct {
	sockets map[uuid.UUID]*socketStruct
	locker  sync.RWMutex
}

/* WebSocketコネクション一覧 */
var melodyManager = &melodyManagerStruct{sockets: make(map[uuid.UUID]*socketStruct)}

/* 部屋作成 */
func HandleCreateRoom(ctx *gin.Context) {
	room, err := room_service.CreateRoom()
	if err != nil {
		// 500
		log.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	melody := newMelody()

	melodyManager.locker.Lock()
	defer melodyManager.locker.Unlock()

	melodyManager.sockets[room.UUID()] = &socketStruct{
		melody: melody,
		room:   room,
	}

	ctx.String(http.StatusCreated, room.UUID().String())
}

/* 部屋取得 */
func HandleGetRoom(ctx *gin.Context) {
	roomId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		// 400 ID形式間違い
		ctx.String(http.StatusBadRequest, "the format of Room ID [%s] is incorrect.", ctx.PostForm("roomId"))
		return
	}
	// 排他処理（書き込み）
	melodyManager.locker.RLock()
	defer melodyManager.locker.RUnlock()

	socket := melodyManager.sockets[roomId]

	socket.locker.RLock()
	defer socket.locker.RUnlock()

	room := socket.room
	res := room_model.Room{Id: room.UUID().String(), Created: room.Created()}
	ctx.JSON(http.StatusOK, res)

}

/* 部屋一覧取得 */
func HandleGetRoomList(ctx *gin.Context) {
	// 排他処理（書き込み）
	melodyManager.locker.RLock()
	defer melodyManager.locker.RUnlock()

	var res = []room_model.Room{}
	for _, v := range melodyManager.sockets {
		v.locker.RLock()
		defer v.locker.RUnlock()

		room := v.room
		model := room_model.Room{Id: room.UUID().String(), Created: room.Created()}

		res = append(res, model)
	}

	ctx.JSON(http.StatusOK, res)
}

/* 部屋参加 */
func HandleConnect(ctx *gin.Context) {
	roomId, err := uuid.Parse(ctx.PostForm("roomId"))
	if err != nil {
		// 400 ID形式間違い
		ctx.String(http.StatusBadRequest, "the format of Room ID [%s] is incorrect.", ctx.PostForm("roomId"))
		return
	}

	melodyManager.locker.Lock()
	defer melodyManager.locker.Unlock()
	socket := melodyManager.sockets[roomId]
	if socket == nil {
		// 400 存在しない部屋ID
		ctx.String(http.StatusBadRequest, "room ID [%s] dose not exists.", roomId.String())
		return
	}

	// アップグレード
	err = socket.melody.HandleRequest(ctx.Writer, ctx.Request)
	if err != nil {
		// 500 ハンドシェイク失敗
		log.Fatal(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
}

/* メロディ作成 */
func newMelody() *melody.Melody {
	m := melody.New()
	handlers := makeMelodyHandler(m)

	m.HandleConnect(handlers.connectHandler)
	m.HandleMessage(handlers.melodyHandler)
	m.HandleDisconnect(handlers.disconnectHandler)

	return m
}

/* メロディハンドラ */
type MelodyHandler struct {
	/* ws開始 */
	connectHandler func(*melody.Session)
	/* wsメッセージ受信 */
	melodyHandler func(*melody.Session, []byte)
	/* ws終了 */
	disconnectHandler func(*melody.Session)
}

/* メロディハンドラ作成 */
func makeMelodyHandler(m *melody.Melody) MelodyHandler {
	res := MelodyHandler{
		connectHandler: func(s *melody.Session) {
			log.Printf("websocket connection open. [session: %#v]\n", s)
		},
		melodyHandler: func(s *melody.Session, msg []byte) {
			m.Broadcast(msg)
		},
		disconnectHandler: func(s *melody.Session) {
			log.Printf("websocket connection close. [session: %#v]\n", s)
		},
	}

	return res
}

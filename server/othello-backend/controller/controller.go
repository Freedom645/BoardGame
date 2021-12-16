package controller

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/Freedom645/BoardGame/controller/middleware"
	"github.com/Freedom645/BoardGame/controller/model"
	"github.com/Freedom645/BoardGame/controller/model/state_model"
	"github.com/Freedom645/BoardGame/domain/enum/player_type"
	"github.com/Freedom645/BoardGame/domain/game"
	"github.com/Freedom645/BoardGame/domain/room"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gopkg.in/olahol/melody.v1"
)

/* 部屋 */
type socketStruct struct {
	melody *melody.Melody
	room   *room.Room
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
	room, err := room.NewRoom()
	if err != nil {
		// 500
		log.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	melody := newMelody(room)

	melodyManager.locker.Lock()
	defer melodyManager.locker.Unlock()

	melodyManager.sockets[room.UUID()] = &socketStruct{
		melody: melody,
		room:   room,
	}

	res := model.Room{Id: room.UUID().String(), Created: room.Created()}
	ctx.JSON(http.StatusCreated, res)
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

	res := model.Of(socket.room)
	ctx.JSON(http.StatusOK, res)
}

/* 部屋一覧取得 */
func HandleGetRoomList(ctx *gin.Context) {
	userInfo, err := middleware.GetUserFromContext(ctx)
	if err != nil {
		ctx.Status(http.StatusUnauthorized)
		return
	}
	log.Printf("userID = %v", userInfo.Uid)

	// 排他処理（書き込み）
	melodyManager.locker.RLock()
	defer melodyManager.locker.RUnlock()

	var res = []model.Room{}
	for _, v := range melodyManager.sockets {
		model := model.Of(v.room)
		res = append(res, *model)
	}

	ctx.JSON(http.StatusOK, res)
}

/* 部屋参加 */
func HandleConnect(ctx *gin.Context) {
	roomId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		// 400 ID形式間違い
		ctx.String(http.StatusBadRequest, "the format of Room ID [%s] is incorrect.", ctx.Param("id"))
		return
	}

	pt, err := player_type.Of(ctx.Query("pt"))
	if err != nil {
		// 400 プレイヤータイプ不明
		ctx.String(http.StatusBadRequest, "%s [%s]", err.Error(), ctx.Query("pt"))
		return
	}

	// TODO デッドロックの危険性
	melodyManager.locker.RLock()
	socket := melodyManager.sockets[roomId]
	melodyManager.locker.RUnlock()

	if socket == nil {
		// 400 存在しない部屋ID
		ctx.String(http.StatusBadRequest, "room ID [%s] dose not exists.", roomId.String())
		return
	}

	// UID取得
	userInfo, _ := middleware.GetUserFromContext(ctx)

	// アップグレード
	var keys = map[string]interface{}{"userInfo": userInfo, "pt": pt}
	err = socket.melody.HandleRequestWithKeys(ctx.Writer, ctx.Request, keys)
	if err != nil {
		// 500 ハンドシェイク失敗
		log.Fatal(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
}

/* メロディ(WebConnection部屋)作成 */
func newMelody(room *room.Room) *melody.Melody {
	m := melody.New()
	handlers := makeMelodyHandler(m, room)

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
func makeMelodyHandler(m *melody.Melody, r *room.Room) MelodyHandler {
	res := MelodyHandler{
		/* セッションコネクション時 */
		connectHandler: func(s *melody.Session) {
			userInfo := s.Keys["userInfo"].(middleware.UserInfo)
			pt := s.Keys["pt"].(player_type.PlayerType)

			log.Printf("websocket connection open. [%s]\n", userInfo.Uid)

			isJoin := r.AddPlayer(userInfo.Uid, userInfo.Name, pt)

			response, err := procResponse(r)
			if err != nil {
				log.Info(err.Error())
				s.Write([]byte(`"invalid request"`))
				return
			}

			d, err := json.Marshal(model.Message{
				Response: response,
			})
			if err != nil {
				log.Error(err)
				return
			}

			if pt == player_type.Player && isJoin {
				m.Broadcast([]byte(d))
			} else {
				s.Write([]byte(d))
			}
		},
		/* メッセージやり取り時 */
		melodyHandler: func(s *melody.Session, msg []byte) {
			var reqMes model.Message
			if err := json.Unmarshal(msg, &reqMes); err != nil {
				s.Write([]byte(`"invalid format"`))
				return
			}

			userInfo := s.Keys["userInfo"].(middleware.UserInfo)
			if err := procRequest(r, userInfo, reqMes.Request); err != nil {
				log.Info(err.Error())
				s.Write([]byte(`"invalid request"`))
				return
			}

			response, err := procResponse(r)
			if err != nil {
				log.Info(err.Error())
				s.Write([]byte(`"invalid request"`))
				return
			}
			sendResponse(m, &response)
		},

		/* セッション離脱時 */
		disconnectHandler: func(s *melody.Session) {
			userInfo := s.Keys["userInfo"].(middleware.UserInfo)
			log.Printf("websocket connection close. [%s]\n", userInfo.Uid)

			response, err := procResponse(r)
			if err != nil {
				log.Error(err)
				return
			}
			if remain := r.RemovePlayer(userInfo.Uid); remain == 0 {
				// 参加者が0の場合は、部屋削除
				melodyManager.locker.Lock()
				defer melodyManager.locker.Unlock()
				// delete(melodyManager.sockets, r.UUID())

			}
			// 通知
			sendResponse(m, &response)
		},
	}

	return res
}

func procRequest(r *room.Room, userInfo middleware.UserInfo, req model.RequestMessage) error {
	switch r.Step() {
	case room.Pending:
		// 対局承認
		isApproved := req.Pending.IsApproved
		if err := r.Approve(userInfo.Uid, isApproved); err != nil {
			return err
		}
	case room.Black:
		fallthrough
	case room.White:
		// 対局中
		pm := req.Game.Point
		if err := r.Put(userInfo.Uid, game.NewPoint(pm.X, pm.Y)); err != nil {
			return err
		}

	case room.GameOver:
		// ゲーム終了時
		isContinue := req.GameOver.IsContinued
		if err := r.Approve(userInfo.Uid, isContinue); err != nil {
			return err
		}
	}

	return nil
}

func procResponse(r *room.Room) (model.ResponseMessage, error) {
	res := model.ResponseMessage{
		Step:    state_model.Of(r.Step()),
		Board:   r.Stones(),
		Owner:   model.ParsePlayer(r.Owner()),
		Players: model.ParsePlayers(r.Players()),
		Turn:    model.TurnOf(r.Turn()),
	}

	return res, nil
}

func sendResponse(m *melody.Melody, response *model.ResponseMessage) {
	d, err := json.Marshal(model.Message{
		Response: *response,
	})
	if err != nil {
		log.Error(err)
		return
	}

	m.Broadcast([]byte(d))
}

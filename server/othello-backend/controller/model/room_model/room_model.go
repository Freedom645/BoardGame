package room_model

import "time"

type Room struct {
	Id      string    `json:"id"`
	Created time.Time `json:"created"`
}

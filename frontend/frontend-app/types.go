package frontend

import "time"

type Response struct {
	Msg       string    `json:"msg"`
	TimeStamp time.Time `json:"timeStamp"`
}

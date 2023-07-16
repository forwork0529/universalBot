package models


type Message struct {
	MessageID int64 `json:"message_id"`
	FromUserID int64 `json:"from,omitempty"`
	Text string `json:"text,omitempty"`
}
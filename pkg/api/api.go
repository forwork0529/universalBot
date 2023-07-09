package api

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sync"
)

type Server struct {
	bot       *tgbotapi.BotAPI
	pipeLine  []Step // list of actions with input messages
	userIdMap sync.Map
}

func New()*Server{
	server := &Server{}
	server.Use()
	server.Use()
	server.Use()
	return server
}

func (s *Server) Run(){

}


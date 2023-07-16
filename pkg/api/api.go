package api

import (
	"botApi/pkg/bot"
	"botApi/pkg/logger"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"os/signal"
	"sync"
	"time"
)

type Server struct {
	bot			bot.Interface
	pipeLine  	[]Step // list of actions with every input message
	userIdMap 	sync.Map  // map wih users which a in computing now
}

// new server with string api
func New(bot bot.Interface)*Server{

	server := &Server{bot : bot}

	server.Use(CheckNoCurrentProcess)
	server.Use(CheckEmptyOrStartMessage)
	server.Use(CheckErrorMessage)
	server.Use(PrepareInputMessage)
	// server.Use(CheckPasswordMessage?)

	// Here is the common string api
	server.Use(func(stepNumber int, server *Server, update tgbotapi.Update){




	})

	return server
}

func (s *Server) Run(){
	logger.Info(`server started...`)

	// Handle interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	<-c

	for name, closer := range a.publicCloser {
		time.Sleep(15 * time.Second)
		logger.Info(ctx, "Stop %s", name)
		closer()
	}
	return nil
}


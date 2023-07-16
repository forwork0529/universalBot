package api

import (
	"botApi/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Step func(stepNumber int, server *Server, update tgbotapi.Update)

func (s *Server) Use(check Step){
	s.pipeLine = append(s.pipeLine, check)
}


// checks that there is no actual process with user which send the message
func CheckNoCurrentProcess(currentStepNumber int, server *Server, update tgbotapi.Update){

	logger.Debug("CheckNoCurrentProcess()..")
	_, ok := server.userIdMap.Load(update.Message.From.ID)
	if ok{
		logger.Debug("CheckNoCurrentProcess(): cant create more then 1 process for one userId")
		return
	}

	server.userIdMap.Store(update.Message.From.ID, struct{}{})
	logger.Debugf("CheckNoCurrentProcess(): %v stored in userIdMap", update.Message.From.ID)

	defer func(server *Server, userId int64){
		server.userIdMap.Delete(userId)
		logger.Debugf("CheckNoCurrentProcess(): %v deleted from userIdMap", update.Message.From.ID)
	}(server, update.Message.From.ID)


	if canDoNextStep(currentStepNumber + 1, len(server.pipeLine)){
		server.pipeLine[currentStepNumber + 1](currentStepNumber, server, update)
	}

}

func RunPipeLine(stepNumber int, server *Server, update tgbotapi.Update){
	server.pipeLine[stepNumber](stepNumber,server, update)
	return
}

func canDoNextStep(stepNumber int, lenList int)bool{
	var can = true
	if stepNumber < 0 || stepNumber >= lenList{
		can = false
	}
	return can
}

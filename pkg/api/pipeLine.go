package api

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Step func(server *Server, update tgbotapi.Update) bool

func (s *Server) Use(check Step){
	s.pipeLine = append(s.pipeLine, check)
}
// Run for the range of the pipeLine list, if logic ends return true
func (s *Server) RunPipeLine(u tgbotapi.Update)bool{
	var end bool
	for _,  step := range s.pipeLine {
		end = step(s, u)
		if end{
			return end
		}
	}
	return end
}

// checks that there is no actual processs with user wich send the message
func CheckIsCurrentProcess(server *Server, update tgbotapi.Update) bool{
	_, ok := server.userIdMap.Load(update.Message.From.ID)
	if ok{
		return true
	}
	server.userIdMap.Store(update.Message.From.ID, struct{}{})

}


func Step1 (stepNumber int, server *Server, update tgbotapi.Update){
	if update.Message.Text == "start"{
		return
	}
	if stepNumber +1 > len(server.pipeLine) -1 {
		return
	}
}


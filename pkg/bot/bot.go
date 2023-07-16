package bot

import (
	"botApi/pkg/config"
	"botApi/pkg/logger"
	"botApi/pkg/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type Interface interface{

}

func New(shutDownChannel <-chan struct{}) (*Bot, error){

	// создаём api бота
	botAPI, err :=  tgbotapi.NewBotAPI(config.EnvVars.TelegramToken)
	if err != nil{
		logger.Errorf("tgbotapi.NewBotAPI(len token: %v): %v", len(config.EnvVars.TelegramToken), err.Error())
		return nil, err
	}
	botDebugStatus, err := strconv.ParseBool(config.EnvVars.BotDebugStatus)
	if err != nil{
		logger.Errorf("strconv.ParseBool(config.EnvVars.BotDebugStatus): %v", err.Error())
		botDebugStatus = false
	}
	botAPI.Debug = botDebugStatus
	logger.Infof("Authorized on account %s", botAPI.Self.UserName)
	logger.Infof("Bot started, debug mode: %v", botAPI.Debug)

	// api бота создан

	// создаём рабочую структуру бота
	bot := &Bot{
		bot : botAPI,
	}

	// создаём канал с сообщениями
	bot.msgChan = bot.GetUpdatesChan()

	// фиксируем в структуре канал завершения работы
	bot.shutDownChannel = shutDownChannel

	return bot, nil
}

type Bot struct{
	bot *tgbotapi.BotAPI
	msgChan <-chan models.Message
	shutDownChannel <-chan struct{}
}

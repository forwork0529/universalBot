package bot

import (
	"botApi/pkg/config"
	"botApi/pkg/logger"
	"botApi/pkg/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"time"
)

// GetUpdatesChan starts and returns a channel for getting updates.
func (bot *Bot) GetUpdatesChan() <-chan models.Message {

	bufferSize, err := strconv.Atoi(config.EnvVars.BotBuffer)
	if err != nil{
		logger.Errorf("strconv.Atoi(config.EnvVars.BotBuffer): %v", err.Error())
		bufferSize = 10
	}
	ch := make(chan models.Message, bufferSize )

	go func() {

		// конфигурируем опрос
		uc := tgbotapi.UpdateConfig{Offset: 0, Limit: 15, Timeout: 60}

		// опрашиваем в цикле
		for {
			select {
			case <-bot.shutDownChannel:
				close(ch)
				logger.Info("updates channel successfully closed")
				return
			default:
			}

			// по конфигурациям получаем обновления от бота
			updates, err := bot.bot.GetUpdates(uc)
			if err != nil {
				logger.Errorf("bot.GetUpdates(): %v, retrying in 3 seconds..." , err.Error())
				time.Sleep(time.Second * 3)
				continue
			}

			// отправляем полученные выше обновления в канал
			for _, update := range updates {
				if update.UpdateID >= uc.Offset {
					uc.Offset = update.UpdateID + 1

					var message models.Message
					message.MessageID = int64(update.UpdateID)
					message.FromUserID = update.Message.From.ID
					message.Text = update.Message.Text

					ch <- message
				}
			}
		}
	}()

	return ch
}

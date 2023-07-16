package api

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"log"
	"strings"
)

func informativeAnswer(bot *tgbotapi.BotAPI, update tgbotapi.Update,  newMsgText, path string)  {

	simpleAnswerText(bot, update, newMsgText)

	fileNames , noFiles := returnFilesInPath(path)

	if noFiles{
		simpleAnswerPhoto(bot, update, Path + "/hangres/hangres.jpg")
		simpleAnswerText(bot, update, "Ассортимент пополняется)..")
	}

	for _, productName := range fileNames{

		simpleAnswerText(bot, update, TrimSuffix(productName))
		simpleAnswerPhoto(bot, update, fmt.Sprintf("%s/%s" ,path, productName))
	}
}

// Функция отправляющая  просто текст сообщения
func simpleAnswerText(bot *tgbotapi.BotAPI, update tgbotapi.Update,  newMsgText string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, newMsgText )
	bot.Send(msg)
}

// Функция отправляющая  просто текст сообщения и выводящая клавиатуру
func simpleAnswerTextWithKeyBoard(bot *tgbotapi.BotAPI, update tgbotapi.Update,  newMsgText string){
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, newMsgText )
	// msg.ReplyToMessageID = update.Message.MessageID
	msg.ReplyMarkup = numericKeyboard
	bot.Send(msg)
}


func simpleAnswerPhoto(bot *tgbotapi.BotAPI, update tgbotapi.Update, photoPath string){
	msgPic := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath(photoPath) )
	_, err := bot.Send(msgPic)
	if err != nil{
		log.Printf("cant send image: %v\n", err.Error())
	}
}

// Функция возвращающая список именён файлов в папке и значение говорящее пуст ли список
func returnFilesInPath(path string)([]string, bool){
	files, err := ioutil.ReadDir(path)
	if err != nil{
		log.Printf("Cant read direcory: %v\n", err.Error())
	}
	var fileNames []string
	for _, fileInfo := range files{
		fileNames = append(fileNames, fileInfo.Name())
	}
	if len(fileNames) < 1{
		return fileNames, true
	}
	return fileNames, false

}

func TrimSuffix(toPrintProductName string )string{
	if idx := strings.IndexByte(toPrintProductName, '.'); idx >= 0 {
		toPrintProductName = toPrintProductName[:idx]
	}
	return toPrintProductName
}

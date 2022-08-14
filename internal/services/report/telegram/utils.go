package telegram

import (
	telegramBot "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

//url to send request: https://api.telegram.org/bot<token>/METHOD_NAME
const (
	url    = "https://api.telegram.org/bot"
	chatId = 1819062414
)

func ReportErrorMessageTelegram(message string) {
	err := godotenv.Load()
	if err != nil {
		logrus.Errorf("Error loading .env file. Cannot read telegram token,detail: %v", err)
		return
	}
	accessToken := os.Getenv("TELEGRAM_ACCESS_TOKEN")

	bot, err := telegramBot.NewBotAPI(accessToken)
	if err != nil {
		logrus.Error("Something trouble with access_token while creating new bot: ", err)
	}
	logrus.Infof("Telegram bot created, detail: %v", bot)
	bot.Debug = true

	msg := telegramBot.NewMessage(chatId, message)
	_, err = bot.Send(msg)
	if err != nil {
		logrus.Errorf("Something went wrong while sending message, detail: %v", err)
		return
	}
	logrus.Info("Message sent successfully")
}

func AutoBot() {
	//@dev lặp code
	err := godotenv.Load()
	if err != nil {
		logrus.Errorf("Error loading .env file. Cannot read telegram token,detail: %v", err)
		return
	}
	accessToken := os.Getenv("TELEGRAM_ACCESS_TOKEN")

	bot, err := telegramBot.NewBotAPI(accessToken)
	if err != nil {
		logrus.Error("Something trouble with access_token while creating new bot: ", err)
	}

	updateCf := telegramBot.NewUpdate(0)
	updateCf.Timeout = 30
	updates, _ := bot.GetUpdatesChan(updateCf)
	for update := range updates {
		if update.Message != nil {
			msg := telegramBot.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}

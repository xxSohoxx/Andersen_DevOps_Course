package main

import (
	"log"
	
	//Importing the package tgbotapi which make it much easier to create a Telegram bot
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	//NewBotAPI creates a new BotAPI instance
	bot, err := tgbotapi.NewBotAPI("YOUR_BOT_TOKEN")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	//NewUpdate gets updates since the last Offset
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//GetUpdatesChan starts and returns a channel for getting updates.
	updates, err := bot.GetUpdatesChan(u)

	//Loop for managing incomming messages
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if update.Message.Text == "Git" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Here is my Git repository: https://github.com/xxSohoxx/Andersen_DevOps_Course")
			bot.Send(msg)
		} else if update.Message.Text == "Tasks" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Here is my list of tasks:\nTask1 - Ansible\nTask2 - Bash\nTask3 - Git\nTask4 - Docker")
			bot.Send(msg)
		} else if update.Message.Text == "Task1" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "https://github.com/xxSohoxx/Andersen_DevOps_Course/tree/main/ansible_task")
			bot.Send(msg)
		} else if update.Message.Text == "Task2" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "https://github.com/xxSohoxx/Andersen_DevOps_Course/tree/main/Bash_task")
			bot.Send(msg)
		} else if update.Message.Text == "Task3" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "https://github.com/xxSohoxx/Andersen_DevOps_Course/tree/main/git_task")
			bot.Send(msg)
		} else if update.Message.Text == "Task4" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "https://github.com/xxSohoxx/Andersen_DevOps_Course/tree/main/docker_task")
			bot.Send(msg)
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Sorry, so far I understand only these commands:\nGit\nTasks\nTask# (where # is number of task)")
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Sorry, so far I understand only these commands:\nGit\nTask\nTask# (where # is number of task)")
		//msg.ReplyToMessageID = update.Message.MessageID

		//bot.Send(msg)
	}
}

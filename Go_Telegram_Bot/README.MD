# About the task #
In this task we need to create a Telegram Bot by using Go programming language. Bot needs to know at least 3 commands:  
* Git - returns link of this Git repository
* Tasks - returns list of completed tasks from the Andersen DevOps Course
* Task# - returns link for repository with completed task, where "#" is a number of task.


## About the bot itself ##  
This bot is created by using non-standard library - `https://github.com/go-telegram-bot-api/telegram-bot-api` and based on polling method.  

Telegram Bot - **@AndersenJuneBot**

## If you want to deploy it in your environment ##
* [Here](https://core.telegram.org/bots/#3-how-do-i-create-a-bot) you can find how to create a Telegram bot and get a Token  
* Then you need to install Go on your working station. You can use the guidence from [here](https://golang.org/doc/install).   
* Please ensure the library is installed and up to date by running *go get -u github.com/go-telegram-bot-api/telegram-bot-api*  
* Clone this repository or just copy the code to your working directory. Don't forget to replace **YOUR_BOT_TOKEN** by your bot token at the row number 13 while we create new BotAPI instance!
* And then run *go build .* command to compile the code into an executable file. And then just run this file. 

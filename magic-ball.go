// create telegram bot
// bot will work as Magic 8 ball
// will randomly answer to question
// from predefined set of answers
package main

import (
	"math/rand"
	"strings"
	"time"

	"github.com/tucnak/telebot"
)

func main() {
	bot, err := telebot.NewBot("137864605:AAGX9Wb_FaNsK9W0RJlsJo7J9RrO_hTFOOo")
	if err != nil {
		return
	}

	messages := make(chan telebot.Message)
	bot.Listen(messages, 1*time.Second)

	for message := range messages {
		if message.Text == "/hi" {
			bot.SendMessage(message.Chat,
				"Hello, "+message.Sender.FirstName+"!", nil)
		}
		if strings.Contains(message.Text, "/?") {
			cmd := strings.Split(message.Text, " ")
			if len(cmd) == 1 {
				bot.SendMessage(message.Chat,
					"no question asked", nil)
				continue
			}

			bot.SendMessage(message.Chat,
				answers[rand.Intn(len(answers)-1)], nil)

		}
	}
}

var answers = []string{
	"It is certain",
	"It is decidedly so",
	"Without a doubt",
	"Yes definitely",
	"You may rely on it",
	"As I see it, yes",
	"Most likely",
	"Outlook good",
	"Yes",
	"Signs point to yes",
	"Reply hazy try again",
	"Ask again later",
	"Better not tell you now",
	"Cannot predict now",
	"Concentrate and ask again",
	"Don't count on it",
	"My reply is no",
	"My sources say no",
	"Outlook not so good",
	"Very doubtful",
}

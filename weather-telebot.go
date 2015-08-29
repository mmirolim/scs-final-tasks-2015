// create telegram bot
// create weather telegram bot
// by using openweather api
// bot should reply what temp and wind speed for asked city 
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/tucnak/telebot"
)

type Weather struct {
	Main Main `json:"main"`
	Wind Wind `json:"wind"`
}
type Main struct {
	Temp float32 `json:"temp"`
}
type Wind struct {
	Speed float32 `json:"speed"`
}

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
		if strings.Contains(message.Text, "/w") {
			cmd := strings.Split(message.Text, " ")
			if len(cmd) == 1 {
				bot.SendMessage(message.Chat,
					"no city chosen", nil)
				continue
			}
			uri := "http://api.openweathermap.org/data/2.5/weather?q=" + cmd[1] + "&units=metric"
			res, err := http.Get(uri)
			if err != nil {
				bot.SendMessage(message.Chat,
					err.Error(), nil)
				continue
			}

			data, err := ioutil.ReadAll(res.Body)
			if err != nil {
				bot.SendMessage(message.Chat,
					err.Error(), nil)
				continue
			}
			var w Weather
			if err := json.Unmarshal(data, &w); err != nil {
				bot.SendMessage(message.Chat,
					err.Error(), nil)
				continue
			}

			bot.SendMessage(message.Chat,
				fmt.Sprintf("Temp is %v and Wind speed is %v in %s", w.Main.Temp, w.Wind.Speed, cmd[1]), nil)

		}
	}
}

package comands

import (
	"HidirAi/ai"
	"fmt"
	"strconv"

	"gopkg.in/telebot.v3"
)

func Start(c telebot.Context) error {
	DeleteUserData(c)

	userid := strconv.FormatInt(c.Sender().ID, 10)
	name := c.Sender().FirstName
	surname := c.Sender().LastName
	username := c.Sender().Username

	user := fmt.Sprintf("Немного информации о пользователе. Его telegramid: %s, его имя: %s, его фамилия(если есть): %s, и его username в телеграмме: %s", userid, name, surname, username)

	hi, err := ai.HidirAi(userid, "Привет", user)

	msg, _ := c.Bot().Send(c.Chat(), "HidirAI это ИИ и он может ошибатся. Проверяйте дважды важную информацию")
	c.Bot().Pin(msg)

	_, err = c.Bot().Raw("sendMessage", map[string]any{
		"chat_id": c.Chat().ID,
		"text":    hi.Text,
		"link_preview_options": map[string]any{
			"url":                "https://github.com/Jasurios/HidirAi/blob/master/menu.png?raw=true",
			"prefer_large_media": true,
			"show_above_text":    true,
		},
	})

	return err
}

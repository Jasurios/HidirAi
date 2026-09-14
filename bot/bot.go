package bot

import (
	"HidirAi/ai"
	"HidirAi/ai/tools"
	"HidirAi/bot/comands"
	"HidirAi/rwfile"
	"encoding/base64"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"gopkg.in/telebot.v3"
)

func Init() error {
	token := os.Getenv("TOKEN")
	settings := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(settings)
	if err != nil {
		return err
	}

	Run(bot)

	log.Println("Bot started")
	bot.Start()

	return nil
}

func Run(bot *telebot.Bot) {
	bot.Handle("/start", func(c telebot.Context) error {
		return comands.Start(c)
	})
	bot.Handle("/delmydata", func(c telebot.Context) error {
		return comands.DeleteUserData(c)
	})

	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		userid := strconv.FormatInt(c.Sender().ID, 10)
		rwfile.CheckSpace(userid)

		// msg, anim := Thinkng(c)

		result, err := ai.HidirAi(userid, c.Text(), "")
		if err != nil {
			log.Println(err)
		}

		// close(anim)

		// switch result.Type {
		// case ai.ResultPhoto:
		// 	c.Bot().Edit(msg, result.Photo)
		// 	return nil
		// case ai.ResultText:
		// 	c.Bot().Edit(msg, result.Text,telebot.ModeMarkdown)
		// 	return nil
		// }

		switch result.Type {
		case ai.ResultPhoto:
			return c.Send(result.Photo)
		case ai.ResultText:
			return c.Send(result.Text, telebot.ModeHTML)
		}

		return c.Send("Ошибка")
	})

	bot.Handle(telebot.OnVoice, func(c telebot.Context) error {
		userid := strconv.FormatInt(c.Sender().ID, 10)
		rwfile.CheckSpace(userid)

		tgFile := &telebot.File{FileID: c.Message().Voice.FileID}

		reader, err := c.Bot().File(tgFile)
		if err != nil {
			log.Println(err)
			return c.Send("Ошибка")
		}
		defer reader.Close()

		audioBytes, err := io.ReadAll(reader)
		if err != nil {
			log.Println(err)
			return c.Send("Ошибка")
		}

		resultAudeo, err := tools.STT(audioBytes)
		if err != nil {
			log.Println(err)
			return c.Send("Ошибка")
		}

		result, err := ai.HidirAi(userid, resultAudeo, "")
		if err != nil {
			log.Println(err)
		}

		switch result.Type {
		case ai.ResultPhoto:
			return c.Send(result.Photo)
		case ai.ResultText:
			return c.Send(result.Text, telebot.ModeMarkdown)
		}

		return c.Send("Ошибка")
	})

	bot.Handle(telebot.OnPhoto, func(c telebot.Context) error {
		userid := strconv.FormatInt(c.Sender().ID, 10)
		rwfile.CheckSpace(userid)

		tgFile := &telebot.File{FileID: c.Message().Photo.FileID}

		reader, _ := c.Bot().File(tgFile)
		defer reader.Close()

		data, _ := io.ReadAll(reader)
		base64Img := base64.StdEncoding.EncodeToString(data)

		text, err := tools.ViewImage(userid, base64Img, "фото")
		if err != nil {
			log.Println(err)
			return c.Send("Ошибка")
		}

		result, err := ai.HidirAi(userid, text, "")
		if err != nil {
			log.Println(err)
		}

		switch result.Type {
		case ai.ResultPhoto:
			return c.Send(result.Photo)
		case ai.ResultText:
			return c.Send(result.Text, telebot.ModeMarkdown)
		}

		return c.Send("Ошибка")
	})

	bot.Handle(telebot.OnSticker, func(c telebot.Context) error {
		userid := strconv.FormatInt(c.Sender().ID, 10)
		rwfile.CheckSpace(userid)

		tgFile := &telebot.File{FileID: c.Message().Sticker.FileID}

		reader, _ := c.Bot().File(tgFile)
		defer reader.Close()

		data, _ := io.ReadAll(reader)
		base64Img := base64.StdEncoding.EncodeToString(data)

		text, err := tools.ViewImage(userid, base64Img, "стикер")
		if err != nil {
			log.Println(err)
			return c.Send("Ошибка")
		}

		result, err := ai.HidirAi(userid, text, "")
		if err != nil {
			log.Println(err)
		}

		switch result.Type {
		case ai.ResultPhoto:
			return c.Send(result.Photo)
		case ai.ResultText:
			return c.Send(result.Text, telebot.ModeMarkdown)
		}

		return c.Send("Ошибка")
	})
}

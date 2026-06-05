package main

import (
	"HidirAi/NukeUserData"
	"HidirAi/ai"
	"HidirAi/limits"
	"HidirAi/skils/VoiceToText"
	"HidirAi/skils/registry"
	"HidirAi/skils/vision"

	"encoding/base64"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

func main() {
	godotenv.Load("./config.env")
	log.Println("Vars loaded")

	token := os.Getenv("TOKEN")

	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, _ := telebot.NewBot(pref)

	bot.Handle("/start", func(c telebot.Context) error {
		name := c.Sender().FirstName
		return c.Send("Hi " + name)
	})

	bot.Handle("/nuke", func(c telebot.Context) error {
		userid := strconv.FormatInt(c.Sender().ID, 10)
		NukeUserData.User(userid, c.Args())

		return c.Send("Success")
	})

	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		userid := strconv.FormatInt(c.Sender().ID, 10)
		NukeUserData.CheckSpace(userid)

		result, err := ai.Askhid(userid, c.Text())
		if err != nil {
			log.Println(err)
		}

		switch result.Type {
		case registry.ResultPhoto:
			return c.Send(result.Photo)
		case registry.ResultText:
			return c.Send(result.Text)
		}

		return c.Send("Ошибка какаято")
	})

	bot.Handle(telebot.OnVoice, func(c telebot.Context) error {
		userid := strconv.FormatInt(c.Sender().ID, 10)
		NukeUserData.CheckSpace(userid)

		tgFile := &telebot.File{FileID: c.Message().Voice.FileID}

		audiobytes, _ := c.Bot().File(tgFile)
		defer audiobytes.Close()

		localFile, _ := os.Create("./users/" + userid + "/lastest.ogg")
		defer localFile.Close()

		io.Copy(localFile, audiobytes)
		log.Printf("Audio saved")

		result, err := VoiceToText.Convert(userid)
		if err != nil {
			log.Println(err)
		}
		
		switch result.Type {
		case registry.ResultPhoto:
			return c.Send(result.Photo)
		case registry.ResultText:
			return c.Send(result.Text)
		}

		return c.Send("Ошибка какаято")
	})

	bot.Handle(telebot.OnPhoto, func(c telebot.Context) error {
		userid := strconv.FormatInt(c.Sender().ID, 10)
		NukeUserData.CheckSpace(userid)

		photo := c.Message().Photo
		tgFile := &telebot.File{FileID: photo.FileID}

		fileReader, _ := c.Bot().File(tgFile)
		defer fileReader.Close()

		data, _ := io.ReadAll(fileReader)
		base64Img := base64.StdEncoding.EncodeToString(data)

		prompt := c.Message().Caption
		if prompt == "" {
			prompt = "Что на этом изображении?"
		}

		result, err := vision.AskHidImage(userid, prompt, base64Img)
		if err != nil {
			log.Println(err)
			return c.Send("Ошибка")
		}

		return c.Send(result)
	})

	limits.Check()
	log.Println("Bot started")

	bot.Start()
}

package tools

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"gopkg.in/telebot.v3"
)

func GenImage(prompt string) *telebot.Photo {
	encodedPrompt := url.QueryEscape(prompt)
	apiURL := fmt.Sprintf("https://image.pollinations.ai/prompt/%s?width=512&height=512&nologo=true", encodedPrompt)

	resp, err := http.Get(apiURL)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(resp.StatusCode)
		return nil
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	photo := &telebot.Photo{
		File: telebot.FromReader(bytes.NewReader(data)),
	}
	photo.File.FileLocal = "image.png"
	return photo
}
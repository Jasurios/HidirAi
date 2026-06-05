package imagegen

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"gopkg.in/telebot.v3"
)

func Gen(prompt string) *telebot.Photo {
	token := os.Getenv("HF_API")

	body, _ := json.Marshal(map[string]interface{}{
		"inputs": prompt,
	})

	req, err := http.NewRequest("POST",
		"https://router.huggingface.co/hf-inference/models/black-forest-labs/FLUX.1-schnell",
		bytes.NewReader(body),
	)
	if err != nil {
		log.Println(err)
		return nil
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		data, _ := io.ReadAll(resp.Body)
		log.Println("HF error:", string(data))
		return nil
	}

	data, _ := io.ReadAll(resp.Body)

	photo := &telebot.Photo{
		File: telebot.FromReader(bytes.NewReader(data)),
	}
	photo.File.FileLocal = "image.png"
	return photo
}
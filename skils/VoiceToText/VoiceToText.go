package VoiceToText

import (
	"context"
	"log"

	// "log"
	"os"

	"HidirAi/ai"
	"HidirAi/skils/registry"

	"github.com/sashabaranov/go-openai"
)

func Convert(userid string)(registry.SkillResult, error){
	config := openai.DefaultConfig(os.Getenv("API"))
	config.BaseURL = os.Getenv("URL")
	client := openai.NewClientWithConfig(config)

	file := "./users/"+userid+"/lastest.ogg"

	req := openai.AudioRequest{
		Model:    "whisper-large-v3",
		FilePath: file,
		Language: "ru",
	}

	resp, err := client.CreateTranscription(context.Background(), req)
	if err != nil {
		log.Println(err)
	}

	answer , _ := ai.Askhid(userid,resp.Text)
	return answer , nil
}

func main() {
	
}
package vision

import (
	"HidirAi/rwfile"
	"context"
	"encoding/json"
	"os"

	"github.com/sashabaranov/go-openai"
)

func AskHidImage(userid, prompt, base64Img string) (string, error) {
	config := openai.DefaultConfig(os.Getenv("API"))
	config.BaseURL = os.Getenv("URL")
	client := openai.NewClientWithConfig(config)

	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: "meta-llama/llama-4-scout-17b-16e-instruct",
		Messages: []openai.ChatCompletionMessage{
			{
				
				Role:    openai.ChatMessageRoleSystem,
				Content: "Не используй markdown форматирование, таблицы и звёздочки. Пиши простым текстом.",
			},
			{
				Role: openai.ChatMessageRoleUser,
				MultiContent: []openai.ChatMessagePart{
					{
						Type: openai.ChatMessagePartTypeImageURL,
						ImageURL: &openai.ChatMessageImageURL{
							URL: "data:image/jpeg;base64," + base64Img,
						},
					},
					{
						Type: openai.ChatMessagePartTypeText,
						Text: prompt,
					},
				},
			},
		},
	})
	if err != nil {
		return "", err
	}

	userHistory := []openai.ChatCompletionMessage{}
	if rwfile.Check(userid) {
		var saved []openai.ChatCompletionMessage
		json.Unmarshal([]byte(rwfile.Read(userid)), &saved)
		if len(saved) > 0 {
			userHistory = saved
		}
	}

	userHistory = append(userHistory, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: "[пользователь отправил фото] "+prompt,
	})
	userHistory = append(userHistory, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: resp.Choices[0].Message.Content,
	})

	historyBytes, _ := json.Marshal(userHistory)
	rwfile.Write(userid, string(historyBytes))

	return resp.Choices[0].Message.Content, nil
}

func main() {

}

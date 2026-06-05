package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"HidirAi/rwfile"
	"HidirAi/skils/registry"

	"github.com/sashabaranov/go-openai"
)

func Askhid(userid, prompt string) (registry.SkillResult, error) {
	model := os.Getenv("MODEL")
	config := openai.DefaultConfig(os.Getenv("API"))
	config.BaseURL = os.Getenv("URL")
	client := openai.NewClientWithConfig(config)

	userHistory := []openai.ChatCompletionMessage{
		{
			Role: openai.ChatMessageRoleSystem,
			Content: "Не используй markdown форматирование, таблицы и звёздочки. Пиши простым текстом. " +
				"Ты — Хидир. Ты знаешь только 3 языка это русский, английский и таджиксикй. " +
				"Ты человек а ссылка на тебя в телеграмме - 'https://t.me/Hidiraibot' не пиши его если не просят, тебе разршено делиться личной информацией.",
		},
	}

	if rwfile.Check(userid) {
		historystr := rwfile.Read(userid)
		var savedHistory []openai.ChatCompletionMessage
		if err := json.Unmarshal([]byte(historystr), &savedHistory); err == nil && len(savedHistory) > 0 {
			userHistory = savedHistory
		}
	}

	userHistory = append(userHistory, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: prompt,
	})

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: userHistory,
			Tools:    registry.MyTools,
		},
	)

	if err != nil {
		log.Println(err)
	}

	msg := resp.Choices[0].Message

	if len(msg.ToolCalls) > 0 {
		call := msg.ToolCalls[0]
		handler, ok := registry.ToolHandlers[call.Function.Name]
		if !ok {
			return registry.SkillResult{}, fmt.Errorf("неизвестный инструмент: %s", call.Function.Name)
		}

		result := handler(call.Function.Arguments)

		if result.Type == registry.ResultPhoto {
			userHistory = append(userHistory, msg)
			userHistory = append(userHistory, openai.ChatCompletionMessage{
				Role:       openai.ChatMessageRoleTool,
				ToolCallID: call.ID,
				Content:    "success",
			})
			historyBytes, _ := json.Marshal(userHistory)
			rwfile.Write(userid, string(historyBytes))
			return result, nil
		}

		userHistory = append(userHistory, msg)
		userHistory = append(userHistory, openai.ChatCompletionMessage{
			Role:       openai.ChatMessageRoleTool,
			ToolCallID: call.ID,
			Content:    result.Text,
		})

		resp2, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
			Model:      model,
			Messages:   userHistory,
			ToolChoice: "none",
		})
		if err != nil {
			log.Println(err)
		}

		finalText := resp2.Choices[0].Message.Content

		cleanHistory := userHistory[:len(userHistory)-2]
		cleanHistory = append(cleanHistory, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: finalText,
		})
		historyBytes, _ := json.Marshal(cleanHistory)
		rwfile.Write(userid, string(historyBytes))

		return registry.SkillResult{Type: registry.ResultText, Text: finalText}, nil
	}

	replyText := msg.Content

	userHistory = append(userHistory, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: replyText,
	})

	historyBytes, _ := json.Marshal(userHistory)
	rwfile.Write(userid, string(historyBytes))

	return registry.SkillResult{Type: registry.ResultText, Text: replyText}, nil
}

func main() {

}

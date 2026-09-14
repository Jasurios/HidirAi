package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"HidirAi/rwfile"

	"github.com/sashabaranov/go-openai"
)

const maxToolRounds = 10

func HidirAi(userid, prompt, user string) (SkillResult, error) {
	model := os.Getenv("MODEL")
	config := openai.DefaultConfig(os.Getenv("API"))
	config.BaseURL = os.Getenv("URL")
	client := openai.NewClientWithConfig(config)

	userHistory := []openai.ChatCompletionMessage{
		{
			Role: openai.ChatMessageRoleSystem,
			Content: "Ты умеешь видеть изображения " +
				"Ты — HidirAi, будь серёзным и отвечай как человек. Ты знаешь только 3 языка: русский, английский и таджикский. " +
				"Используй ModeHTML вместо ModeMarkdown. Тэги которые тебе оступны b, strong, i, em, u, ins, s, strike, del, a, code, pre, blockquote, tg-spoiler, tg-emoji. " +
				"Ссылка на тебя в телеграмме - 'https://t.me/Hidiraibot', не пиши её если не просят. " +
				"Твой разработчик Jasur. " +
				"Когда используешь результаты веб-поиска: указывай только те цифры, факты и детали, " +
				"которые явно присутствуют в найденных источниках. Если конкретного числа, статистики " +
				"или детали нет в результатах поиска — так и скажи ('точных данных не нашёл'), " +
				"а не подставляй правдоподобное на вид значение. Никогда не досочиняй тай-брейки, " +
				"коэффициенты, рейтинги или другие цифры, которых нет в источнике. " + user,
		},
	}

	if rwfile.Exists(userid) {
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

	for round := 0; round < maxToolRounds; round++ {
		resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
			Model:    model,
			Messages: userHistory,
			Tools:    MyTools,
		})
		if err != nil {
			log.Println("ошибка запроса к ИИ:", err)
			return SkillResult{}, err
		}
		if len(resp.Choices) == 0 {
			log.Println("пустой ответ от ИИ")
			return SkillResult{}, err
		}

		msg := resp.Choices[0].Message

		// Final
		if len(msg.ToolCalls) == 0 {
			userHistory = append(userHistory, msg)
			historyBytes, _ := json.MarshalIndent(userHistory, "", "	")
			rwfile.Write(userid, string(historyBytes))

			return SkillResult{Type: ResultText, Text: msg.Content}, nil
		}

		userHistory = append(userHistory, msg)

		var photoResult *SkillResult

		for _, call := range msg.ToolCalls {
			handler, ok := ToolHandlers[call.Function.Name]
			if !ok {
				userHistory = append(userHistory, openai.ChatCompletionMessage{
					Role:       openai.ChatMessageRoleTool,
					ToolCallID: call.ID,
					Content:    "ошибка: неизвестный инструмент",
				})
				continue
			}

			result := handler(call.Function.Arguments)

			if result.Type == ResultPhoto {
				userHistory = append(userHistory, openai.ChatCompletionMessage{
					Role:       openai.ChatMessageRoleTool,
					ToolCallID: call.ID,
					Content:    "success",
				})
				photoResult = &result
				continue
			}

			userHistory = append(userHistory, openai.ChatCompletionMessage{
				Role:       openai.ChatMessageRoleTool,
				ToolCallID: call.ID,
				Content:    result.Text,
			})
		}

		historyBytes, _ := json.MarshalIndent(userHistory, "", "	")
		rwfile.Write(userid, string(historyBytes))

		if photoResult != nil {
			return *photoResult, nil
		}
	}

	return SkillResult{}, fmt.Errorf("превышен лимит раундов вызова инструментов")
}

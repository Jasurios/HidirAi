package rwfile

import (
	"encoding/json"
	"strconv"

	"os"

	"github.com/sashabaranov/go-openai"
)

func CheckSpace(userid string) {
	limit, _ := strconv.Atoi(os.Getenv("MESSAGE"))
	file := "./users/" + userid + "/" + userid + ".json"
	history, _ := os.ReadFile(file)

	var messages []openai.ChatCompletionMessage
	json.Unmarshal(history, &messages)

	if len(messages) > limit {
		cut := 2
		for cut < len(messages) && messages[cut].Role != "user" {
			cut++
		}

		messages = append(messages[:1], messages[cut:]...)

		jsonBytes, _ := json.MarshalIndent(messages, "", "    ")
		os.WriteFile(file, jsonBytes, 0644)
	}
}
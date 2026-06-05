package NukeUserData

import (
	"encoding/json"
	"strconv"

	// "log"
	"os"

	"github.com/sashabaranov/go-openai"
)

func User(userid string, args []string) {
	admin := os.Getenv("ADMIN")
	if len(args) >= 1 {
		if userid == admin && args[0] == "all" {
			os.RemoveAll("./users/")
		} else if userid == admin {
			os.RemoveAll("./users/" + args[0])
		}
	} else {
		os.RemoveAll("./users/" + userid)
	}
}

func CheckSpace(userid string) {
	message := os.Getenv("MESSAGE")
	intm, _ := strconv.Atoi(message)
	file := "./users/" + userid + "/" + userid + ".json"
	history, _ := os.ReadFile(file)

	var messages []openai.ChatCompletionMessage
	json.Unmarshal(history, &messages)

	if len(messages) > intm {
		deleteCount := 2
		if len(messages) > 2 && messages[3].Role == "tool" {
			deleteCount = 3
		}

		messages = append(messages[0:1], messages[1+deleteCount:]...)

		jsonBytes, _ := json.MarshalIndent(messages, "", "    ")
		os.WriteFile(file, jsonBytes, 0644)
	}
}

func main() {

}

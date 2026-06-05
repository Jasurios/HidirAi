package limits

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"fmt"
)

func Check() {
	API := os.Getenv("API")
    model := os.Getenv("MODEL")

	mod := fmt.Sprintf(`{
        "model": "%s",
        "messages": [{"role": "user", "content": "."}]
    }`,model)

    jsonBody := []byte(mod)

    req, _ := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(jsonBody))
    req.Header.Add("Authorization", "Bearer "+API)
    req.Header.Add("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println(err)
        return
    }
    defer resp.Body.Close()

    log.Println("----------limits----------")
    log.Printf("Осталось запросов на этот день (RPD): %s\n", resp.Header.Get("X-Ratelimit-Remaining-Requests"))
    log.Printf("Осталось токенов на эту минуту (TPM): %s\n", resp.Header.Get("X-Ratelimit-Remaining-Tokens"))
    log.Printf("Сброс минутного лимита токенов: %s\n", resp.Header.Get("X-Ratelimit-Reset-Tokens"))
    log.Printf("Сброс дневного лимита запросов: %s\n", resp.Header.Get("X-Ratelimit-Reset-Requests"))
}
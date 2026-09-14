package bot

import (
	"math/rand"
	"strings"
	"time"

	"gopkg.in/telebot.v3"
)

func Thinkng(c telebot.Context) (*telebot.Message, chan struct{}) {
	rand.Int()
	thinkingMsg, _ := c.Bot().Send(c.Chat(), "Думаю")

	anim := make(chan struct{})

	go func() {
		i := 0
		for {
			select {
			case <-anim:
				return
			case <-time.After(700 * time.Millisecond):
				dots := strings.Repeat(".", i)
				c.Bot().Edit(thinkingMsg, "Думаю"+dots)
				i++
			}
		}
	}()

	return thinkingMsg, anim
}

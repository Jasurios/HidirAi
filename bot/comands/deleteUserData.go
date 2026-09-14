package comands

import (
	"HidirAi/rwfile"
	"fmt"

	"gopkg.in/telebot.v3"
)

func DeleteUserData(c telebot.Context) error {
	userid := fmt.Sprintf("%d",c.Sender().ID)

	args := c.Args()
	if len(args) > 0{
		rwfile.DeleteUserData(userid, args[0])
	}else{
		rwfile.DeleteUserData(userid, "")
	}


	return c.Send("Готово")
}
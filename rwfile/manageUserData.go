package rwfile

import (
	"log"
	"os"
)

func Read(userid string) string {
	filename := userid

	history , _ := os.ReadFile("./users/"+filename+"/"+filename+".json")
	return string(history)
}

func Write(userid, history string){
	filename := userid

	err := os.Mkdir("./users/"+filename, 0755)
	if err != nil{
		os.Mkdir("./users/", 0755)
		os.Mkdir("./users/"+filename, 0755)
	}
	os.WriteFile("./users/"+filename+"/"+filename+".json", []byte(history), 0644)
}

func Exists(userid string) bool {
	filename := userid

	_, err := os.Stat("./users/"+filename+"/"+filename+".json")
	return err == nil
}

func DeleteUserData(userid, args string) error {
	admin := os.Getenv("ADMIN")

	log.Println("Пользователь удалил свои данные")

	if args != "" {
		if userid == admin && args == "all" {
			err := os.RemoveAll("./users/")
			if err != nil{
				return err
			}
		} else if userid == admin {
			err := os.RemoveAll("./users/" + args)
			if err != nil{
				return err
			}
		}
	} else {
		err := os.RemoveAll("./users/" + userid)
		if err != nil{
			return err
		}
	}
	return nil
}
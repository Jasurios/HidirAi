package main

import (
	"HidirAi/bot"
	"log"

	logger "github.com/Jasurios/SimpleLogger"
	"github.com/joho/godotenv"
)

func main() {
	err := logger.Start("server.log")
	if err != nil{
		log.Fatalln(err)
	}
	log.Println("Logger started")

	err = godotenv.Load("./config.env")
	if err != nil{
		log.Fatalln(err)
	}
	log.Println("Vars loaded")

	err = bot.Init()
	if err != nil{
		log.Fatalln(err)
	}
}
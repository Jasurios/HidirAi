package qrcode

import (
	"bytes"

	"github.com/skip2/go-qrcode"
	"gopkg.in/telebot.v3"
)



func Create(text string)*telebot.Photo{
	qrBytes, _ := qrcode.Encode(text, qrcode.Low, 256)

	photo := &telebot.Photo{
		File: telebot.FromReader(bytes.NewReader(qrBytes)),
	}

	photo.File.FileLocal = "qrcode.png"
	
	photo.Caption = "QRcode " + text

	return photo
}


func main() {
	
}
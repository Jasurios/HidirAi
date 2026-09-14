package tools

import (
	"bytes"

	"github.com/skip2/go-qrcode"
	"gopkg.in/telebot.v3"
)

func MakeQrcode(text string) *telebot.Photo {
	qrBytes, _ := qrcode.Encode(text, qrcode.Low, 256)

	photo := &telebot.Photo{
		File: telebot.FromReader(bytes.NewReader(qrBytes)),
	}

	photo.Caption = "QRcode: " + text

	return photo
}

package weather

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"net/http"

	"gopkg.in/telebot.v3"
)

func GetWeatherIn(city, lang string) *telebot.Photo {
	url := fmt.Sprintf("https://wttr.in/%s.png?lang=%s", city, lang)

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		log.Println(err)
		return nil
	}

	bounds := img.Bounds()
	cropped := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(0, 0, bounds.Max.X, bounds.Max.Y-14))

	var buf bytes.Buffer
	png.Encode(&buf, cropped)

	photo := &telebot.Photo{
		File: telebot.FromReader(bytes.NewReader(buf.Bytes())),
	}
	photo.File.FileLocal = "weather.png"
	photo.Caption = "Погода в " + city

	return photo
}
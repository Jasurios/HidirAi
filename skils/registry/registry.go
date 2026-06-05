package registry

import (
	"encoding/json"

	"HidirAi/skils/imagegen"
	"HidirAi/skils/qrcode"
	"HidirAi/skils/search"
	"HidirAi/skils/weather"

	"github.com/sashabaranov/go-openai"
	"gopkg.in/telebot.v3"
)

type ResultType string

const (
	ResultText  ResultType = "text"
	ResultPhoto ResultType = "photo"
)

type SkillResult struct {
	Type  ResultType
	Text  string
	Photo *telebot.Photo
}

var MyTools = []openai.Tool{
	{
		Type: openai.ToolTypeFunction,
		Function: &openai.FunctionDefinition{
			Name:        "create_qrcode",
			Description: "Создать qr код",
			Parameters: json.RawMessage(`{
				"type": "object",
				"properties": {
					"text": {"type": "string", "description": "Ссылка или текст для QR кода. Если пользователь не дал текст то попроси его что добавить в QR код"}
				},
				"required": ["text"]
			}`),
		},
	},
	{
		Type: openai.ToolTypeFunction,
		Function: &openai.FunctionDefinition{
			Name:        "get_weather",
			Description: "Поиск погоды на 3 дня",
			Parameters: json.RawMessage(`{
				"type": "object",
				"properties": {
					"city": {"type": "string", "description": "Город в котором живёт пользователь или просто город в котором он хочет узнать погоду. Если пользователь не сказал где он живёт или не дал город то попроси у него"},
					"lang": {"type": "string", "description": "Код языка на которм пишет пользователь например: таджикский и русский-ru, английский-en"}
				},
				"required": ["city"]
			}`),
		},
	},
	{
		Type: openai.ToolTypeFunction,
		Function: &openai.FunctionDefinition{
			Name:        "web_search",
			Description: "Поиск информации в интернете. Используй для новостей, актуальных данных, цен, людей и всего что может измениться со временем. Если пользователь спрашивает что-то чего ты не знаешь — ищи.",
			Parameters: json.RawMessage(`{
				"type": "object",
				"properties": {
					"query": {"type": "string", "description": "Поисковый запрос. Пиши запрос на том же языке на котором пишет пользователь"}
				},
				"required": ["query"]
			}`),
		},
	},
	{
		Type: openai.ToolTypeFunction,
		Function: &openai.FunctionDefinition{
			Name:        "image_gen",
			Description: "Генерация картинки по описанию. Prompt ВСЕГДА переводи на английский язык перед передачей, даже если пользователь написал на русском или таджикском.",
			Parameters: json.RawMessage(`{
				"type": "object",
				"properties": {
					"prompt": {"type": "string", "description": "То что пользователь хочет видить в изрбражении"}
				},
				"required": ["prompt"]
			}`),
		},
	},
}

var ToolHandlers = map[string]func(args string) SkillResult{
	"get_weather": func(args string) SkillResult {
		var a struct {
			City string `json:"city"`
			Lang string `json:"lang"`
		}
		json.Unmarshal([]byte(args), &a)
		photo := weather.GetWeatherIn(a.City, a.Lang)
		if photo == nil {
			return SkillResult{Type: ResultText, Text: "Не удалось получить погоду"}
		}
		return SkillResult{Type: ResultPhoto, Photo: photo}
	},
	"image_gen": func(args string) SkillResult {
		var a struct {
			Prompt string `json:"prompt"`
		}
		json.Unmarshal([]byte(args), &a)
		photo := imagegen.Gen(a.Prompt)
		if photo == nil {
			return SkillResult{Type: ResultText, Text: "Не удалось сгенерировать картинку"}
		}
		return SkillResult{Type: ResultPhoto, Photo: photo}
	},
	"create_qrcode": func(args string) SkillResult {
		var a struct {
			Text string `json:"text"`
		}
		json.Unmarshal([]byte(args), &a)
		return SkillResult{Type: ResultPhoto, Photo: qrcode.Create(a.Text)}
	},
	"web_search": func(args string) SkillResult {
		var a struct {
			Query string `json:"query"`
		}
		json.Unmarshal([]byte(args), &a)
		return SkillResult{Type: ResultText, Text: search.Search(a.Query)}
	},
}

package ai

import (
	"HidirAi/ai/tools"
	"encoding/json"

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
				"required": ["city","lang"]
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
	// {
	// 	Type: openai.ToolTypeFunction,
	// 	Function: &openai.FunctionDefinition{
	// 		Name:        "call_phone",
	// 		Description: "Звонок за таджикские номера. Используй если пользователь хочет чтобы ты позвонил по какому-то номеру. Например: закзать еду, узнать цену, забронировать отель. Отвечай как человек при разговоре с опонентом, сбрасывай трубку только тогда когда услышишь характерные досвидание, пока или же когда понимаешь что разговор окончен",
	// 		Parameters: json.RawMessage(`{
	// 			"type": "object",
	// 			"properties": {
	// 				"number": {"type": "string", "description": "Номер телефона на который хочет позвонить пользователь"}
	// 			},
	// 			"required": ["number"]
	// 		}`),
	// 	},
	// },
}

var ToolHandlers = map[string]func(args string) SkillResult{
	"get_weather": func(args string) SkillResult {
		var a struct {
			City string `json:"city"`
			Lang string `json:"lang"`
		}
		json.Unmarshal([]byte(args), &a)
		photo := tools.GetWeatherIn(a.City, a.Lang)
		if photo == nil {
			return SkillResult{Type: ResultText, Text: "Не удалось получить погоду"}
		}
		return SkillResult{Type: ResultPhoto, Photo: photo}
	},
	"create_qrcode": func(args string) SkillResult {
		var a struct {
			Text string `json:"text"`
		}
		json.Unmarshal([]byte(args), &a)
		return SkillResult{Type: ResultPhoto, Photo: tools.MakeQrcode(a.Text)}
	},
	"web_search": func(args string) SkillResult {
		var a struct {
			Query string `json:"query"`
		}
		json.Unmarshal([]byte(args), &a)

		text, err := tools.SearchInWeb(a.Query)
		if err != nil {
			text = err.Error()
		}
		if text == "" {
			text = "По этому запросу ничего не найдено."
		}

		return SkillResult{Type: ResultText, Text: text}
	},
	"image_gen": func(args string) SkillResult {
		var a struct {
			Prompt string `json:"prompt"`
		}
		json.Unmarshal([]byte(args), &a)
		photo := tools.GenImage(a.Prompt)
		if photo == nil {
			return SkillResult{Type: ResultText, Text: "Не удалось сгенерировать картинку"}
		}
		return SkillResult{Type: ResultPhoto, Photo: photo}
	},
	// "call_phone": func(args string) SkillResult {
	// 	var a struct {
	// 		Number string `json:"number"`
	// 	}
	// 	json.Unmarshal([]byte(args), &a)
	// 	return SkillResult{Type: ResultText, Text: "Вы разговариваите с человеком пожалуйста поздаровайтесь дождитесь ответа опонента и скажите для чего вы ему позвонили а затем чтобы завершить диалог напишите stop_call"}
	// 	// return SkillResult{Type: ResultText, Text: phone.Call(a.Number)}
	// },
}

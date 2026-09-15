# HidirAi

[![Go Version](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go&logoColor=white)](go.mod)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![Telegram Bot](https://img.shields.io/badge/Telegram-Bot-26A5E4?logo=telegram&logoColor=white)](https://t.me/Hidiraibot)
[![Made with telebot.v3](https://img.shields.io/badge/made%20with-telebot.v3-informational)](https://github.com/tucnak/telebot)

**A modular, multimodal Telegram AI assistant written in Go.**


HidirAi understands text, voice, photos, and stickers, and can call tools — web search, weather, QR code generation, image generation — to answer with real, grounded information instead of guessing. Every user gets their own persistent conversation history.

## Features

- 💬 **Text conversations** with full history, stored per user and replayed on every message
- 🎙 **Voice messages** — transcribed via Groq Whisper (STT) and answered like any text message
- 🖼 **Vision** — understands photos and stickers sent to the bot
- 🛠 **Tool calling**, with each tool as a small, independent module:
  - `web_search` — for anything time-sensitive (news, prices, current facts)
  - `get_weather` — 3-day forecast for a given city
  - `create_qrcode` — generates a QR code from text/link
  - `image_gen` — generates an image from a description (prompt auto-translated to English)
- 🧭 **Grounded answers** — the system prompt explicitly forbids inventing numbers, stats, or facts not present in search results
- 🌐 **Multilingual** — responds in Russian, English, or Tajik depending on the user
- 🗑 **Privacy command** (`/delmydata`) — users can wipe their own stored history
- 📦 **Automatic history trimming** (`rwfile.CheckSpace`) to keep per-user storage bounded
- 📞 *(work in progress, currently disabled)* — a `call_phone` tool for placing real phone calls

## Project Structure

```
HidirAi/
├── main.go              # entrypoint: loads env, starts logger, starts bot
├── bot/
│   ├── bot.go            # Telegram handlers: text, voice, photo, sticker
│   ├── animation.go       # animated "thinking..." status message
│   └── comands/           # /start, /delmydata
├── ai/
│   ├── ai.go              # core loop: sends messages to the model, resolves tool calls
│   ├── toolsConfig.go      # tool definitions + handlers registry
│   └── tools/              # STT, vision, web search, weather, QR code, image gen
└── rwfile/                # per-user history read/write + space management
```

## Tech Stack

- **Language:** Go
- **Bot API:** [telebot.v3](https://github.com/tucnak/telebot)
- **LLM client:** [go-openai](https://github.com/sashabaranov/go-openai) (OpenAI-compatible endpoint, model/URL configurable via env)
- **Speech-to-text & vision:** Groq
- **QR codes:** [go-qrcode](https://github.com/skip2/go-qrcode)
- **Logging:** [SimpleLogger](https://github.com/Jasurios/SimpleLogger)

## Getting Started

You have two options:

- **Build it yourself** — clone the repo and set it up with Go (below).
- **Download a ready-to-run build** — check the [Releases](https://github.com/Jasurios/HidirAi/releases) page, where a prebuilt file is provided.

### Build from source

```bash
git clone https://github.com/Jasurios/HidirAi.git
cd HidirAi
go mod download
cp config.env.example config.env
```

Fill in `config.env`:

```env
# ---------- Telegram ----------
TOKEN=
ADMIN=

# ---------- AI APIs ----------
TAVILY_API=

GROQ_API=
GROQ_URL=
MODEL_VISION=
MODEL_SPEECH=

API=
URL=
MODEL=

# ---------- Settings ----------
MESSAGE=16
```

Then run:

```bash
go run main.go
```

## How it works

Every incoming message (text, transcribed voice, or a description of a photo/sticker) goes through `ai.HidirAi`, which:
1. loads the user's saved history (or starts a fresh one with the system prompt),
2. sends it to the model along with the available tools,
3. if the model calls a tool, runs it and feeds the result back — repeating up to 10 rounds,
4. saves the updated history and returns the final text or photo to send back.

## License

Apache License 2.0 — see [LICENSE](LICENSE) for details.

## Author

Built by [Jasurios](https://github.com/Jasurios).
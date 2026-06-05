<div align="center">

# 🤖 HidirAi

**Умный Telegram-бот на Go с поддержкой ИИ, голоса и зрения**

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue?style=for-the-badge)](LICENSE)
[![Telegram](https://img.shields.io/badge/Telegram-Bot-26A5E4?style=for-the-badge&logo=telegram&logoColor=white)](https://t.me)
[![Author](https://img.shields.io/badge/Author-Jasurios-FF6B6B?style=for-the-badge)](https://github.com/Jasurios)

</div>

---

## ✨ Возможности

| Фича | Описание |
|------|----------|
| 💬 **Текстовый чат** | Общение с ИИ-моделью через обычные сообщения |
| 🎤 **Голосовые сообщения** | Автоматическое распознавание речи (Voice-to-Text) |
| 👁️ **Анализ изображений** | Отправь фото — бот опишет что на нём |
| 🌐 **Веб-поиск** | Поиск актуальной информации через Tavily API |
| 🧠 **Память контекста** | Помнит историю диалога (настраивается) |
| 🗑️ **NukeUserData** | Команда `/nuke` для полной очистки данных пользователя |
| ⚡ **Лимиты** | Встроенная система контроля нагрузки |

---

## 🏗️ Структура проекта

```
HidirAi/
├── ai/               # Основная логика запросов к ИИ
├── limits/           # Контроль лимитов и нагрузки
├── NukeUserData/     # Очистка пользовательских данных
├── rwfile/           # Работа с файловой системой
├── skils/
│   ├── registry/     # Реестр типов ответов
│   ├── VoiceToText/  # Распознавание голоса
│   └── vision/       # Анализ изображений
├── main.go           # Точка входа
├── config.env.example
├── go.mod
└── LICENSE
```

---

## 🚀 Быстрый старт

### 1. Клонируй репозиторий

```bash
git clone https://github.com/Jasurios/HidirAi.git
cd HidirAi
```

### 2. Настрой конфиг

```bash
cp config.env.example config.env
```

Заполни `config.env`:

```env
# --- Telegram ---
TOKEN=your_telegram_bot_token
ADMIN=your_telegram_user_id

# --- AI ---
TAVILY_API=your_tavily_api_key       # для веб-поиска
API=your_groq_api_key                # или другой провайдер
URL=https://api.groq.com/openai/v1  # URL API (можно заменить)
HF_API=your_huggingface_api_key      # для голоса/изображений

# --- Модели ---
MODEL=openai/gpt-oss-120b            # можно поменять на любую

# --- Настройки ---
MESSAGE=32                           # количество сообщений в памяти
```

### 3. Установи зависимости и запусти

```bash
go mod download
go run main.go
```

---

## 🧪 Команды бота

| Команда | Описание |
|---------|----------|
| `/start` | Приветствие — бот поздоровается по имени |
| `/nuke` | Полная очистка всех данных пользователя |
| _любой текст_ | Ответ от ИИ с учётом контекста диалога |
| _голосовое_ | Распознавание речи + ответ |
| _фото_ | Анализ изображения (можно с подписью-промптом) |

---

## 📦 Зависимости

- [`gopkg.in/telebot.v3`](https://github.com/tucnak/telebot) — Telegram Bot API
- [`github.com/joho/godotenv`](https://github.com/joho/godotenv) — загрузка `.env` конфига
- [Groq API](https://groq.com) / совместимые OpenAI-эндпоинты — языковые модели
- [Hugging Face API](https://huggingface.co) — speech-to-text и vision
- [Tavily API](https://tavily.com) — веб-поиск в реальном времени

---

## 📄 Лицензия

Этот проект распространяется под лицензией **Apache License 2.0**.

### Что это значит для тебя?

✅ **Можно:**
- Использовать в личных и коммерческих проектах
- Свободно изменять и распространять код
- Встраивать в другие продукты (в т.ч. закрытые)
- Делать приватные форки

❌ **Нельзя:**
- Использовать имя автора (**Jasurios**) для продвижения производных продуктов без разрешения
- Снимать с себя ответственность — лицензия предоставляется «как есть», без гарантий
- Убирать уведомление об авторских правах и текст лицензии из исходников

⚠️ **Обязательно:**
- При распространении (исходников или бинарников) — включить копию лицензии (`LICENSE`)
- Указать все изменения, внесённые в оригинальный код
- Сохранить указание на оригинального автора: **Jasurios**

> Полный текст лицензии: [`LICENSE`](LICENSE)  
> Официальный текст: [apache.org/licenses/LICENSE-2.0](https://www.apache.org/licenses/LICENSE-2.0)

---

## 👤 Автор

Создано с ❤️ и **[Jasurios](https://github.com/Jasurios)**

---

<div align="center">

Если проект полезен — поставь ⭐ на GitHub, это реально помогает!

</div>

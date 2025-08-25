# Backend проекта _jp-song-trainer_

Этот раздел содержит бекэнд-документацию проекта: структуру, запуск и статус разработки.

---

## 📌 Описание

Бэкенд для проекта **«Японский по песням»** (jp-song-trainer).  
Написан на **Go (Chi)**, предоставляет REST API для перевода песен, комментариев и истории переводов.  
На текущем этапе реализован **MVP-каркас** с мок-данными.

---

## 📂 Структура

```txt
└── 📁backend
    └── 📁cmd
        └── 📁server
            ├── main.go
    └── 📁internal
        └── 📁handlers
            ├── comments.go
            ├── history.go
            ├── translate.go
        └── 📁models
        └── 📁repo
        └── 📁services
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── README.md
    └── server
```

---

## 🚀 Запуск локально

Требуется Go версии **1.24+**.

```bash
cd backend
go run ./cmd/server
```

Сервер будет доступен по адресу:
`http://localhost:8080`

Примеры запросов:

```bash
# Перевод
curl -X POST http://localhost:8080/api/translate \
  -H "Content-Type: application/json" \
  -d '{"title":"My Song","artist":"Some Artist","lyrics":"君の名は..."}'

# История
curl http://localhost:8080/api/history
```

---

## 🐳 Запуск через Docker

Собрать образ:

```bash
docker build -t jp-backend .
```

Запустить контейнер:

```bash
docker run -p 8080:8080 jp-backend
```

Проверить:

```bash
curl http://localhost:8080/api/history
```

---

## 📖 Документация

- [📄 Specification](/docs/specification.md) - функциональные/нефункциональные требования, архитектура, API.
- [📄 Vision & Scope](/docs/vision_and_scope.md) - цели и границы проекта.
- [📄 User Stories](/docs/user_stories.md) - сценарии использования и Acceptance Criteria.
- [📄 Glossary](/docs/glossary.md) - термины и определения.
- [📄 Swagger API](/docs/api/swagger.yaml) - OpenAPI спецификация.

---

## ✅ Статус разработки (MVP Backend)

- Реализован каркас на Go (Chi).
- Подняты эндпоинты:
  - `POST /api/translate` — перевод текста (мок).
  - `GET /api/history` — список сохранённых переводов (мок).
  - `DELETE /api/history/{id}` — удаление из истории (мок).
  - `PUT /api/comments/{id}` — обновление комментариев (мок).
- Собран рабочий Docker-образ, эндпоинты протестированы локально и через контейнер.

---

## 🔮 Ближайшие шаги

- Подключить SQLite для хранения истории переводов.
- Реализовать интеграцию с DeepSeek API (перевод/комментарии).
- Добавить обработку ошибок и коды согласно спецификации.
- Подготовить `docker-compose` для совместного запуска frontend + backend.

---

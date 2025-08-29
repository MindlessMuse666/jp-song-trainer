# 📑 Backend ![MIT](https://img.shields.io/badge/License-MIT-yellow.svg)

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white) ![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white) ![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white) ![Visual Studio Code](https://img.shields.io/badge/Visual%20Studio%20Code-0078d7.svg?style=for-the-badge&logo=visual-studio-code&logoColor=white)

Добро пожаловать в backend-спецификацию сервиса для изучения японского языка через разбор песен и текстов!

---

## Технологический стек

- Go 1.24.6+
- Chi Router
- PostgreSQL
- JWT аутентификация
- Docker & Docker Compose
- Zap для логирования

---

## Быстрый старт

1. Склонируйте репозиторий
2. Создайте файл `.env`:

    ```bash
    cp .env .env
    ```

3. Задайте в созданном окружении `.env` переменные окружения:

    ```txt
    PORT=<YOUR_PORT>
    POSTGRES_USER=<YOUR_POSTGRES_USER>
    POSTGRES_PASSWORD=<YOUR_POSTGRES_PASSWORD>
    POSTGRES_DB=<YOUR_POSTGRES_DB_NAME>
    DATABASE_URL=<YOUR_POSTGRES_URL>
    JWT_SECRET=<YOUR_JWT_SECRET>
    ```

4. Запустите приложение:

    ```bash
    docker-compose up --build
    ```

Сервер будет доступен на `http://localhost:8080`

---

## API Endpoints

### Аутентификация

- `POST /api/auth/register` - Регистрация пользователя
- `POST /api/auth/login` - Вход пользователя
- `GET /api/auth/profile` - Получение профиля (требует JWT)

### Системные

- `GET /health` - Проверка здоровья сервера

---

## Доступ к базам данных

**PostgreSQL**: `localhost:5432`
**pgAdmin**: `http://localhost:5050`

---

## Тестирование API

### Использование curl

**Тест healthcheck**:

```bash
curl http://localhost:8080/health
```

**Регистрация пользователя**:

```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com", "password":"password123"}'
```

**Вход пользователя**:

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com", "password":"password123"}'
```

**Получение профиля (с JWT токеном)**:

```bash
# Скопируйте токен из ответа на запрос входа
curl -X GET http://localhost:8080/api/auth/profile \
  -H "Authorization: Bearer <your-jwt-token>"
```

### Использование Postman

1. Установите Postman: `https://www.postman.com/downloads/`
2. Создайте новый запрос для каждого эндпоинта
3. Для POST-запросов установите:
   - Method: POST
   - Body -> raw -> JSON
   - Введите JSON данные (например):

      ```json
      {
        "email":"test@example.com",
        "password":"password123"
      }
      ```

4. Для защищенных роутов добавьте заголовок:

- Key: Authorization
- Value: Bearer \<your-jwt-token\>

### Примеры запросов

**Успешная регистрация**:

```json
{
  "email": "test@example.com",
  "password": "password123"
}
```

**Успешный ответ**:

```json
{
  "user": {
    "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "email": "test@example.com",
    "role": "user",
    "created_at": "2023-01-01T00:00:00Z"
  },
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

---

## Миграции

Миграции базы данных выполняются автоматически при запуске приложения через Docker Compose. Файлы миграций находятся в папке [migrations/](./migrations/)

---

## Разработка

### Структура проекта

```md
└── 📁backend
    └── 📁cmd
        └── 📁server
            ├── main.go
    └── 📁internal
        └── 📁app
        └── 📁config                  # Конфигурация приложения
            ├── config.go
        └── 📁handler                 # HTTP обработчики
            ├── auth_handler.go
            ├── health.go
        └── 📁middleware              # Промежуточное ПО
            ├── auth.go
        └── 📁model                   # Модели данных
            ├── models.go
        └── 📁repository              # Работа с БД
            ├── user_repository.go
        └── 📁service                 # Бизнес-логика
            ├── auth_service.go
    └── 📁migrations                  # Миграции БД
        ├── 001_init_schema.down.sql
        ├── 001_init_schema.up.sql
    └── 📁pkg
        └── 📁database                # Подключение к БД
            ├── database.go
        └── 📁logger                  # Логирование
            ├── logger.go
    ├── .env
    ├── docker-compose.yml
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    └── README.md
```

### Добавление новых эндпоинтов

1. Создайте обработчик в [internal/handler/](./internal/handler/)
2. Добавьте соответствующий метод в сервис (internal/service/)
3. Добавьте методы работы с БД в репозиторий (internal/repository/)
4. Зарегистрируйте роут в [cmd/server/main.go](./cmd/server/main.go)

---

## Логирование

Приложение использует структурированное логирование через **Zap**. Логи выводятся в консоль в формате JSON.

---

## Troubleshooting

1. **Ошибка подключения к БД**: Убедитесь, что PostgreSQL запущен и переменные окружения настроены правильно
2. **Ошибка миграций**: Проверьте, что файлы миграций находятся в папке [migrations/](./migrations/)
3. **Ошибка JWT**: Убедитесь, что задана переменная `JWT_SECRET` в `.env`

---

## Деплой

Для продакшен деплоя:

- Обновите переменные окружения для продакшена
- Используйте надежный `JWT_SECRET`
- Настройте обратный прокси (`Nginx`)
- Настройте `SSL-сертификаты`

---

### Лицензия

Этот проект распространяется под лицензией MIT - смотри файл [LICENSE](../LICENCE) для деталей.

---

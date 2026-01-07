# Notely (Backend)
Notely — это современная веб-платформа, где пользователи могут создавать красиво оформленные публикации с кастомными цветовыми схемами, делиться ими с сообществом и находить интересный контент через систему подписок.

## Tech Stack
- [Go](https://go.dev/), [Gin](https://github.com/gin-gonic/gin), [GORM](https://gorm.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [godotenv](https://github.com/joho/godotenv)

## Installation
Клонируем репозиторий:
```bash
  git clone https://github.com/maltira/Notely_Backend.git
  cd Notely_Backend
```
Устанавливаем зависимости:
```bash
  go mod tidy
```
.env:
```bash
  DB_HOST=<host>
  DB_PASS=<pass>
  DB_USER=<user>
  DB_NAME=<name>
  DB_PORT=<port>
  APP_HOST=<host>
  APP_PORT=<port>
  SECRET=<jwt-secret>
```
Запускаем сервер:
```bash
  go run cmd/main.go
```

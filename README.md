# trade-platform

# Итоговое задание летней школы 2020. Направление Backend
“Торговая площадка продажи игровых ключей”

Сервис, позволяющий пользователям покупать и продавать игровые ключи для определенной игровой площадки.

## Доступ через домен
- https://calm-headland-56829.herokuapp.com

## Сборка через Docker

### Пулл образа из Docker Hub:
```docker run -it -p 8080:8080 ialmostdeveloper/trade-platform```

### Сборка из Dockerfile:
```docker build -t trade-platform . ```

```docker run -it -p 8080:8080 trade-platform```

## Сборка ручками:
### Зависимости для сборки:
- github.com/mattn/go-sqlite3
- github.com/dgrijalva/jwt-go
- github.com/go-openapi/runtime/middleware
- github.com/go-redis/redis
- github.com/google/uuid
- github.com/gorilla/handlers
- github.com/gorilla/mux

### Сборка и запуск:
```go build && ./trade-platform```

## Документация
- http://localhost:8080/docs
- https://calm-headland-56829.herokuapp.com/docs

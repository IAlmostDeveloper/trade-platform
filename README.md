# trade-platform

# Итоговое задание летней школы 2020. Направление Backend
“Торговая площадка продажи игровых ключей”

Сервис, позволяющий пользователям покупать и продавать игровые ключи для определенной игровой площадки.

## Доступ через домен
- https://ff61ebbc434e.ngrok.io

## Сборка через Docker

### Пулл образа из Docker Hub:
```docker pull ialmostdeveloper/trade-platform```

### Сборка из Dockerfile:
```docker build -t trade-platform . ```

### Запуск контейнера:
```docker run -it -e PORT=8080 -e TRADE_PLATFORM_SMTP_PASSWORD='password' -p 8080:8080 ialmostdeveloper/trade-platform```

## Сборка ручками:
### Зависимости для сборки:
- github.com/mattn/go-sqlite3
- github.com/dgrijalva/jwt-go
- github.com/go-openapi/runtime/middleware
- github.com/go-redis/redis
- github.com/google/uuid
- github.com/gorilla/handlers
- github.com/gorilla/mux
- github.com/streadway/amqp

### Также необходимо:
- Redis
- RabbitMQ
- SMTP клиент

### Настройка переменных окружения:
```PORT=8080 && TRADE_PLATFORM_SMTP_PASSWORD="YOUR_PASSWORD"```

### Сборка и запуск:
```go build && ./trade-platform```

## Документация
- http://localhost:8080/docs
- https://ff61ebbc434e.ngrok.io/docs

### Краткое руководство по использованию (Примеры использования эндпоинтов - см. документацию)
- Для начала работы пользователю необходимо заругистрироваться (эндпоинт /register), затем пройти аутентификацию (эндпоинт /auth). 
- Разделение на продавцов и покупателей отсутствует, пользователь может и продавать ключи, и покупать их.
#### Покупка продукта
- Для покупки какого либо ключа продукта необходимо узнать, какие продукты размещены на платформе (эндпоинт GET /products). 
- Для покупки ключа продуктанеобходимо ввести название этого продукта, и получить идентификатор, по которому выдается ключ после оплаты (эндпоинт /purchase). 
- Далее создается платеж с указанным идентификатором, и пользователь получает идентификатор платежной сессии (эндпоинт POST /payment). 
- Затем пользователю необходимо ввести данные банковской карты и идентификатор платежной сессии (эндпоинт /validate).
- После успешной валидации карты и оплаты покупки пользователь получает ключ на свою почту, указанную при регистрации, продавец получает оповещение с информацией о покупке на свой веб-сервер.
#### Загрузка продукта
- Для загрузки продукта пользователь вводит его информацию (имя продукта, ключ продукта, цену, коммиссию), используя эндпоинт POST /products. 
- При продаже продукта пользователь получает на веб-сервер, указанный при регистрации, информацию о покупателе и купленном продукте.

FROM golang



ADD . /go/src/trade-platform
RUN go get -t github.com/mattn/go-sqlite3
RUN go get -t github.com/dgrijalva/jwt-go
RUN go get -t github.com/go-openapi/runtime/middleware
RUN go get -t github.com/go-redis/redis
RUN go get -t github.com/google/uuid
RUN go get -t github.com/gorilla/handlers
RUN go get -t github.com/gorilla/mux
RUN go get -t github.com/streadway/amqp

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 
COPY ./entrypoint.sh /entrypoint.sh
RUN apt-get update && apt-get install -y redis rabbitmq-server
ENTRYPOINT [ "/app/entrypoint.sh" ]

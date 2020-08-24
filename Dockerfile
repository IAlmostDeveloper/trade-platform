FROM golang

ADD . /go/src/trade-platform
RUN go get -t github.com/mattn/go-sqlite3
RUN go get -t github.com/dgrijalva/jwt-go
RUN go get -t github.com/go-openapi/runtime/middleware
RUN go get -t github.com/go-redis/redis
RUN go get -t github.com/google/uuid
RUN go get -t github.com/gorilla/handlers
RUN go get -t github.com/gorilla/mux

ENTRYPOINT cd src/trade-platform/ && go run main.go


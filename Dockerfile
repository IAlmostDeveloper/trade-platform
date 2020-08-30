FROM golang

ADD . /go/src/trade-platform
RUN go get -t github.com/mattn/go-sqlite3
RUN go get -t github.com/dgrijalva/jwt-go
RUN go get -t github.com/go-openapi/runtime/middleware
RUN go get -t github.com/go-redis/redis
RUN go get -t github.com/google/uuid
RUN go get -t github.com/gorilla/handlers
RUN go get -t github.com/gorilla/mux
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 
CMD ["/app/main"]


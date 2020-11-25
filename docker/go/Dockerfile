FROM golang:1.15.2

ENV GO111MODULE=on

RUN ["apt-get", "update"]
RUN ["apt-get", "install", "-y", "vim"]

# COPY src/ /go/src/github.com/yuutan0518/go-demo

# WORKDIR /go/src/app
# WORKDIR /go/src/github.com/yuutan0518/go-demo/src

RUN go get -u github.com/cosmtrek/air
RUN go get -u github.com/go-delve/delve/cmd/dlv
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/gin-gonic/gin
RUN go get go.uber.org/zap

CMD air -c .air.toml

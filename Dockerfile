FROM golang:alpine

WORKDIR /app

COPY . .


RUN go get github.com/mmcdole/gofeed
RUN go get github.com/microcosm-cc/bluemonday
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/joho/godotenv

EXPOSE 8020

CMD ["go","run", "main.go"]








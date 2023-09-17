FROM golang:1.20-alpine

ENV ROOT_PATH=/var/www/blue-api

RUN mkdir -p $ROOT_PATH

WORKDIR $ROOT_PATH

COPY . .

RUN go mod tidy

RUN go build main.go

EXPOSE 3001

CMD go run ./main.go


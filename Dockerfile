FROM golang:1.18

RUN apt update

ENV ROOT_PATH=/var/www/blueapi

RUN mkdir -p $ROOT_PATH

WORKDIR $ROOT_PATH

COPY . .

RUN go mod tidy

RUN go build main.go

EXPOSE 3001

CMD go run ./main.go

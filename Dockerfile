FROM golang:1.20

WORKDIR /app

COPY . ./

RUN go mod tidy

RUN go build main.go

EXPOSE 3000

CMD go run ./main.go

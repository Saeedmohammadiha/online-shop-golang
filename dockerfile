FROM golang:1.22.6-alpine

RUN go install github.com/githubnemo/CompileDaemon@latest

WORKDIR /app

COPY go.mod go.sum ./


RUN go mod download

COPY . .


EXPOSE 5000
CMD ["CompileDaemon", "--build=main.go", "--command=./main"]



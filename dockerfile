FROM golang:1.22.6-alpine

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod go.sum ./


RUN go mod download

COPY . .


EXPOSE 5000
CMD ["air"]



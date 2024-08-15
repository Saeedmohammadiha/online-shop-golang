FROM golang:1.22.6-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Install Delve for debugging
RUN apk add git && \
    go install github.com/go-delve/delve/cmd/dlv@latest

EXPOSE 5000 40000

CMD ["dlv", "debug", "--headless", "--listen=:40000", "--api-version=2", "--accept-multiclient", "--log", "/app/main.go"]

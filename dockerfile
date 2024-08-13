FROM golang:1.22.6-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main.exe .

# Set correct permissions for the executable
RUN chmod +x /app/main.exe



EXPOSE 5000

CMD [ "/app/main.exe"]
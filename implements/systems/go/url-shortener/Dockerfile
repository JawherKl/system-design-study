FROM golang:1.24-alpine
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main ./cmd/main.go
CMD ["./main"]
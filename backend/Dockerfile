FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . . .
RUN CGO_ENABLED=0 go build -o main .

# Need curl to perform healthchecks
FROM alpine/curl:latest
WORKDIR /app
COPY configuration/ ./configuration/
COPY migrations/ ./migrations/
COPY --from=builder /app/main .
CMD ["./main"]

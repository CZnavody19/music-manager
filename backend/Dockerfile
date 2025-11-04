FROM golang:1.24

WORKDIR /app
COPY go.mod go.sum  ./
RUN go mod download
RUN go install github.com/air-verse/air@v1.62.0
COPY . .

CMD ["air"]

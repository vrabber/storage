FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod download

ENTRYPOINT ["air", "-c", "deploy/dev/air/air.toml"]
FROM golang:1.24 AS buider

WORKDIR /app

COPY todo-app/go.mod todo-app/go.sum ./
RUN go mod download
RUN go env -w CGO_ENABLED=0

COPY todo-app /app
RUN go build -o . ./cmd/main.go

FROM ubuntu
COPY --from=buider /app app/config/dev.toml app/config/.env usr/local/bin/
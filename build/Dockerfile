FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ../configs/.devcontainer .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main ./

COPY config ./config

COPY Infrastucture ./Infrastucture

ENTRYPOINT ["/app/main"]

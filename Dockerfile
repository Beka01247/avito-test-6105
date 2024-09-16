FROM golang:1.23.0-alpine AS builder

WORKDIR /app

ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0 

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN ls -la

RUN go build -o main ./cmd/main.go

RUN chmod a+x ./main

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

RUN chmod +x main

RUN cat main

EXPOSE 8080

RUN pwd && ls -l

RUN chmod a+x ./main

ENTRYPOINT ["./main"]
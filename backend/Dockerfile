FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server

FROM alpine:latest

RUN apk --no-cache add ca-certificates postgresql-client

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/wait-for.sh .
COPY --from=builder /app/migrations ./migrations

RUN chmod +x wait-for.sh

EXPOSE 8080

CMD ["./wait-for.sh", "db:5432", "--", "./main"]
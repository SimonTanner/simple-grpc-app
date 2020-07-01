FROM golang:1.13 AS builder

WORKDIR /app

ADD . .

WORKDIR /app/services/frontend

RUN go mod download

RUN go test -race ./...

RUN CGO_ENABLED=0 go build -o frontend

FROM alpine

RUN apk add --update --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/services/frontend/frontend /app/

EXPOSE 8080

ENTRYPOINT ./frontend 
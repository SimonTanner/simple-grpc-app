FROM golang:1.13 AS builder

WORKDIR /app

ADD . .

RUN ls -la

WORKDIR /app/frontend/

RUN go mod download

RUN CGO_ENABLED=0 go build -o frontend

RUN ls -la

FROM alpine

RUN apk add --update --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/frontend/frontend /app/

EXPOSE 8080

ENTRYPOINT ./frontend 
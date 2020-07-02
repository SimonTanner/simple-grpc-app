FROM golang:1.13 AS builder

WORKDIR /app

ADD . .

RUN ls -la

WORKDIR /app/backend

RUN go mod download

RUN CGO_ENABLED=0 go build

RUN ls -la

FROM alpine

RUN apk add --update --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/backend/backend /app/

COPY --from=builder /app/db/wait-for.sh /app/

RUN chmod +x ./wait-for.sh

EXPOSE 8090

ENTRYPOINT ./backend 
FROM golang:1.13 AS builder

WORKDIR /app

ADD . .

RUN ls -la

WORKDIR /app

RUN go mod download

RUN CGO_ENABLED=0 go build

RUN ls -la

FROM alpine

RUN apk add --update --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/backend /app/

EXPOSE 8090

ENTRYPOINT ./api 
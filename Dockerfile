FROM golang:1.22.2-alpine AS builder

ARG TOKEN_GITEA
ARG GOPRIVATE

COPY . /app/src/
WORKDIR /app/src/

ENV GOPRIVATE=${GOPRIVATE}

RUN apk add git && git config --global url."https://${TOKEN_GITEA}:x-oauth-basic@gitea.24example.ru/".insteadOf "https://gitea.24example.ru/"

RUN go mod download
RUN go build -o ./bin/service cmd/server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/src/.env .
COPY --from=builder /app/src/bin/service  .

CMD ["./service"]
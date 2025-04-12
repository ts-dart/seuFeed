FROM golang:1.22-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/scraper
RUN go build -o /out/scraper

WORKDIR /app/api
RUN go build -o /out/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /out/scraper ./scraper
COPY --from=builder /out/api ./api

EXPOSE 8080

CMD ["./api"]

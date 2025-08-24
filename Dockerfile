FROM golang:1.24-alpine AS build
RUN apk add --no-cache git gcc musl-dev sqlite-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bot main.go


FROM alpine:latest
RUN apk add --no-cache sqlite-libs
WORKDIR /app
COPY --from=build /app/bot .
COPY .env .env
RUN chmod +x ./bot

CMD ["./bot"]

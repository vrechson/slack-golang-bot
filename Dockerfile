FROM golang:1.14-alpine AS build-env

WORKDIR /app

COPY . .

RUN go mod download

RUN go build

RUN apk add --no-cache udev ttf-freefont bash chromium chromium-chromedriver

EXPOSE 80

ENTRYPOINT ["./bot"]
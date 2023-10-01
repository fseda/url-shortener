FROM golang:latest as base

WORKDIR /usr/app

COPY . .

RUN go build -o app

FROM base AS production

EXPOSE 8080

CMD ["./app"]